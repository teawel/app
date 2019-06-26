package app

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/lists"
	"github.com/teawel/app/types"
	"github.com/teawel/app/utils"
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type HTTPServer struct {
	wel     *Wel
	address string
}

func NewHTTPServer(wel *Wel, address string) *HTTPServer {
	return &HTTPServer{
		wel:     wel,
		address: address,
	}
}

func (this *HTTPServer) Start() error {
	// start fetching
	this.startFetching(60 * time.Second)

	// start web server
	mux := http.NewServeMux()
	mux.HandleFunc("/", this.handleIndex)

	utils.Println("start listening " + this.address)
	server := &http.Server{
		Addr:    this.address,
		Handler: mux,
	}
	err := server.ListenAndServe()
	return err
}

// start fetching
func (this *HTTPServer) startFetching(duration time.Duration) {
	go func() {
		ticker := time.NewTicker(duration)
		for range ticker.C {
			instances, err := this.retrieveInstances()
			if err != nil {
				log.Println("[error]" + err.Error())
				return
			}
			for _, instance := range instances {
				result, err := this.wel.Fetch(instance["options"].(map[string]string))
				if err != nil {
					log.Println("[error]" + err.Error())
					continue
				}

				db, err := dbs.NewDB(Root + "/web/instances/" + this.wel.Id + "/" + instance["id"].(string) + ".db")
				if err != nil {
					log.Println("[error]" + err.Error())
					continue
				}

				err = db.Write(result)
				if err != nil {
					log.Println("[error]" + err.Error())
				}
				_ = db.Close()
			}
		}
	}()
}

// "/"
func (this *HTTPServer) handleIndex(writer http.ResponseWriter, request *http.Request) {
	path := request.URL.Path

	// api
	if strings.HasPrefix(path, "/api/") || path == "/api" {
		this.handleAPI(writer, request)
		return
	}

	// from ajax
	root := Root + "/web"
	if path == "" || path == "/" {
		this.writeFile(writer, root+"/index.html")
	} else if path == "/addSubmit" {
		this.handleAddSubmit(writer, request)
	} else if path == "/instances" {
		this.handleInstances(writer, request)
	} else if path == "/instance" {
		this.handleInstance(writer, request)
	} else if path == "/instance/update" {
		this.handleInstanceUpdate(writer, request)
	} else if path == "/dashboard" {
		this.handleDashboard(writer, request)
	} else {
		if !strings.Contains(filepath.Base(path), ".") {
			path += ".html"
		}
		this.writeFile(writer, root+path)
	}
}

// "/api"
func (this *HTTPServer) handleAPI(writer http.ResponseWriter, request *http.Request) {
	requestPath := strings.TrimPrefix(request.URL.Path, "/api")

	args := []string{}
	for k, v := range request.URL.Query() {
		for _, v1 := range v {
			args = append(args, "-"+k+"="+v1)
		}
	}

	// from ajax
	isFromAjax := request.Header.Get("X-Requested-With") == "XMLHttpRequest"

	if lists.ContainsString([]string{"/version", "/info", "/options", "/all", "/fetch"}, requestPath) {
		w := bytes.NewBuffer([]byte{})
		err := this.wel.RunCmd(requestPath[1:], args, w)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
		} else {
			if !lists.ContainsString([]string{"/version"}, requestPath) {
				writer.Header().Set("Content-Type", "application/json; charset=utf-8")
			}

			// from ajax
			if isFromAjax {
				writer.Write([]byte(`{
	"code": 200,
	"data":`))
			}

			writer.Write(w.Bytes())

			// from ajax
			if isFromAjax {
				writer.Write([]byte(`
}`))
			}
		}

		return
	}

	if strings.HasPrefix(requestPath, "/run/") {
		code := requestPath[len("/run/"):]
		newArgs := append([]string{code}, args...)
		w := bytes.NewBuffer([]byte{})
		err := this.wel.RunCmd("run", newArgs, w)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte(err.Error()))
		} else {
			if isFromAjax {
				writer.Header().Set("Content-Type", "application/json; charset=utf-8")
				writer.Write([]byte(`{
	"code": 200,
	"data": { "output":` + string(utils.JSONEncode(string(w.Bytes()))) + ` }
}`))
			} else {
				writer.Write(w.Bytes())
			}
		}
		return
	}

	writer.Write([]byte(`API Usage:
/version 
   wel version

/info
   wel information

/options
   wel instance options

/all
   wel all information, options, templates

/fetch?option1=value1&...
   fetch values

/run/OPERATION?option1=value1&...
   run instance operation
`))
}

// write response string
func (this *HTTPServer) writeString(writer http.ResponseWriter, s string, contentType string) {
	writer.Header().Set("Content-Type", contentType)
	writer.Write([]byte(s))
}

// write response error
func (this *HTTPServer) writeError(writer http.ResponseWriter, err error) {
	writer.Write([]byte("Error: " + err.Error()))
}

// write response json
func (this *HTTPServer) writeJSON(writer http.ResponseWriter, value interface{}) {
	data, err := json.Marshal(value)
	if err != nil {
		this.writeError(writer, err)
		return
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Write(data)
}

// write response file
func (this *HTTPServer) writeFile(writer http.ResponseWriter, path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		this.writeError(writer, err)
		return
	}

	mimeType := mime.TypeByExtension(filepath.Ext(path))
	if len(mimeType) > 0 {
		writer.Header().Set("Content-Type", mimeType)
	}

	writer.Write(data)
}

func (this *HTTPServer) writeResponse(writer http.ResponseWriter, code int, message string, value interface{}) {
	this.writeJSON(writer, map[string]interface{}{
		"code":    code,
		"message": message,
		"data":    value,
	})
}

func (this *HTTPServer) handleAddSubmit(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseMultipartForm(2 << 32)
	if err != nil {
		this.writeResponse(writer, 400, "ERROR: "+err.Error(), nil)
		return
	}

	name := request.PostFormValue("name")
	if len(name) == 0 {
		this.writeResponse(writer, 400, "Please enter the name", nil)
		return
	}

	interval := request.PostFormValue("interval")
	if len(interval) == 0 {
		this.writeResponse(writer, 400, "Please enter the refresh interval", nil)
		return
	}

	intervalSeconds, err := strconv.ParseInt(interval, 10, 32)
	if err != nil {
		this.writeResponse(writer, 400, "Refresh interval should be a valid digit number", nil)
		return
	}

	options := map[string]string{}
	for _, option := range this.wel.Options {
		basic := option.BasicOption()
		code := basic.Code
		value, skip, err := option.ApplyRequest(request)
		if err != nil {
			this.writeResponse(writer, 400, "ERROR: "+err.Error(), nil)
			return
		}
		if skip {
			continue
		}

		// TODO apply isRequired/validateCode/initCode

		options[code] = types.String(value)
	}

	instance := NewInstance()
	instance.Name = name
	instance.Interval = int(intervalSeconds)
	instance.Options = options
	instance.Dashboards = this.wel.DashboardTemplates
	instance.Operations = this.wel.Operations
	instance.Charts = this.wel.ChartTemplates
	instance.Thresholds = this.wel.ThresholdTemplates

	if len(this.wel.Id) == 0 {
		this.writeResponse(writer, 400, "Wel Id should not be empty", nil)
		return
	}

	if !regexp.MustCompile("^[a-zA-Z0-9.@-]+$").MatchString(this.wel.Id) {
		this.writeResponse(writer, 400, "Wel Id '"+this.wel.Id+"' is invalid", nil)
		return
	}

	instanceId := utils.TimeFormat("YmdHis")
	instance.Id = instanceId
	err = instance.Write(Root + "/web/instances/" + this.wel.Id + "/" + instanceId + ".yml")
	if err != nil {
		this.writeResponse(writer, 400, "ERROR: "+err.Error(), nil)
		return
	}

	this.writeResponse(writer, 200, "", map[string]interface{}{
		"instanceId": instanceId,
	})
}

func (this *HTTPServer) handleInstances(writer http.ResponseWriter, request *http.Request) {
	if len(this.wel.Id) == 0 {
		this.writeResponse(writer, 400, "Wel Id should not be empty", nil)
		return
	}

	instances, err := this.retrieveInstances()
	if err != nil {
		this.writeResponse(writer, 400, err.Error(), nil)
	}

	this.writeResponse(writer, 200, "", map[string]interface{}{
		"instances": instances,
	})
}

func (this *HTTPServer) handleInstance(writer http.ResponseWriter, request *http.Request) {
	instanceId := request.URL.Query().Get("instance")

	instance, err := LoadInstanceFromPath(Root + "/web/instances/" + this.wel.Id + "/" + instanceId + ".yml")
	if err != nil {
		this.writeResponse(writer, 400, "ERROR: "+err.Error(), nil)
		return
	}

	if len(instance.Id) == 0 {
		instance.Id = instanceId
	}

	this.writeResponse(writer, 200, "", instance)
}

func (this *HTTPServer) handleInstanceUpdate(writer http.ResponseWriter, request *http.Request) {
	instanceId := request.URL.Query().Get("instance")

	instance, err := LoadInstanceFromPath(Root + "/web/instances/" + this.wel.Id + "/" + instanceId + ".yml")
	if err != nil {
		this.writeResponse(writer, 400, err.Error(), nil)
		return
	}

	instance.Dashboards = this.wel.DashboardTemplates
	instance.Operations = this.wel.Operations
	instance.Charts = this.wel.ChartTemplates
	instance.Thresholds = this.wel.ThresholdTemplates
	err = instance.Write(Root + "/web/instances/" + this.wel.Id + "/" + instanceId + ".yml")
	if err != nil {
		this.writeResponse(writer, 400, err.Error(), nil)
		return
	}

	this.writeResponse(writer, 200, "", nil)
}

func (this *HTTPServer) handleDashboard(writer http.ResponseWriter, request *http.Request) {
	instanceId := request.URL.Query().Get("instance")
	dashboardId := request.URL.Query().Get("dashboard")
	if len(instanceId) == 0 {
		this.writeResponse(writer, 400, "'instance' parameter should not be empty", nil)
		return
	}
	if len(dashboardId) == 0 {
		this.writeResponse(writer, 400, "'dashboard' parameter should not be empty", nil)
		return
	}

	instance, err := LoadInstanceFromPath(Root + "/web/instances/" + this.wel.Id + "/" + instanceId + ".yml")
	if err != nil {
		this.writeResponse(writer, 400, err.Error(), nil)
		return
	}

	dashboard := instance.FindDashboard(dashboardId)
	if dashboard == nil {
		this.writeResponse(writer, 400, "can not find dashboard with id '"+dashboardId+"'", nil)
		return
	}

	db, err := dbs.NewDB(Root + "/web/instances/" + this.wel.Id + "/" + instance.Id + ".db")
	if err != nil {
		this.writeResponse(writer, 400, err.Error(), nil)
		return
	}
	defer db.Close()

	for _, chart := range dashboard.Charts {
		err := chart.Fetch(db)
		if err != nil {
			log.Println("[error]" + err.Error())
		}
	}

	this.writeResponse(writer, 200, "", dashboard)
}

func (this *HTTPServer) retrieveInstances() (instances []map[string]interface{}, err error) {
	if !regexp.MustCompile("^[a-zA-Z0-9.@-]+$").MatchString(this.wel.Id) {
		err = errors.New("Wel Id '" + this.wel.Id + "' is invalid")
		return
	}

	dir := Root + "/web/instances/" + this.wel.Id + "/"
	result, err := filepath.Glob(dir + "/*.yml")
	if err != nil {
		return
	}

	instances = []map[string]interface{}{}
	for _, file := range result {
		filename := filepath.Base(file)
		if !regexp.MustCompile(`^\d+.yml$`).MatchString(filename) {
			log.Println("[error]skip '" + filename + "': invalid filename")
			continue
		}
		instanceId := filename[:len(filename)-4]
		instance, err := LoadInstanceFromPath(file)
		if err != nil {
			log.Println("[error]'" + filename + "': " + err.Error())
			continue
		}
		if len(instance.Id) == 0 {
			instance.Id = instanceId
		}
		instances = append(instances, map[string]interface{}{
			"id":      instance.Id,
			"name":    instance.Name,
			"options": instance.Options,
		})
	}
	return instances, nil
}
