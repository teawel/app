package app

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-yaml/yaml"
	"github.com/teawel/app/lists"
	"github.com/teawel/app/options"
	"github.com/teawel/app/utils"
	"github.com/vmihailenco/msgpack"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

type Wel struct {
	Id          string `yaml:"id" json:"id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	Developer   string `yaml:"developer" json:"developer"`
	Version     string `yaml:"version" json:"version"`
	Site        string `yaml:"site" json:"site"`

	Options []options.OptionInterface `yaml:"options" json:"options"`

	Apps       []*App       `yaml:"apps" json:"apps"`
	Operations []*Operation `yaml:"operations" json:"operations"`

	ThresholdTemplates []*Threshold   `yaml:"thresholdTemplates" json:"thresholdTemplates"`
	DashboardTemplates []*Dashboard   `yaml:"dashboardTemplates" json:"dashboardTemplates"`
	ChartTemplates     []*ChartCanvas `yaml:"chartTemplates" json:"chartTemplates"`

	fetcher Fetcher
}

func NewWel() *Wel {
	return &Wel{
		Apps:               []*App{},
		Operations:         []*Operation{},
		ThresholdTemplates: []*Threshold{},
		DashboardTemplates: []*Dashboard{},
		ChartTemplates:     []*ChartCanvas{},
		Options:            []options.OptionInterface{},
	}
}

func (this *Wel) AddOption(option ...options.OptionInterface) {
	this.Options = append(this.Options, option...)
}

func (this *Wel) AddApp(app ...*App) {
	this.Apps = append(this.Apps, app...)
}

func (this *Wel) AddOperation(operation ...*Operation) {
	this.Operations = append(this.Operations, operation...)
}

func (this *Wel) FindOperation(code string) *Operation {
	for _, op := range this.Operations {
		if op.Code == code {
			return op
		}
	}
	return nil
}

func (this *Wel) AddThreshold(thresholdTemplate *Threshold) {
	this.ThresholdTemplates = append(this.ThresholdTemplates, thresholdTemplate)
}

func (this *Wel) AddChartCanvas(canvas *ChartCanvas) {
	this.ChartTemplates = append(this.ChartTemplates, canvas)
}

func (this *Wel) FindChart(chartId string) *ChartCanvas {
	for _, chart := range this.ChartTemplates {
		if chart.Id == chartId {
			return chart
		}
	}
	return nil
}

func (this *Wel) AddDashboard(dashboard *Dashboard) {
	this.DashboardTemplates = append(this.DashboardTemplates, dashboard)
}

func (this *Wel) AddChartsToDashboard(dashboard *Dashboard, chartId ...string) {
	if len(chartId) == 0 {
		return
	}
	for _, chartId2 := range chartId {
		chart := this.FindChart(chartId2)
		if chart == nil {
			log.Println("[error]can not find chart with id '" + chartId2 + "'")
			continue
		}
		dashboard.AddChart(chart)
	}
}

func (this *Wel) AddAllChartsToDashboard(dashboard *Dashboard) {
	for _, chart := range this.ChartTemplates {
		dashboard.AddChart(chart)
	}
}

func (this *Wel) OnFetch(fetcher Fetcher) {
	this.fetcher = fetcher
}

func (this *Wel) Fetch(options map[string]string) (result map[string]string, err error) {
	if this.fetcher == nil {
		return nil, errors.New("fetcher should not be nil")
	}

	return this.fetcher(options)
}

func (this *Wel) Run() {
	if len(os.Args) == 0 {
		return
	}

	if len(os.Args) == 1 {
		this.RunCmd("-h", []string{}, os.Stdout)
		return
	}

	if len(os.Args) > 2 {
		err := this.RunCmd(os.Args[1], os.Args[2:], os.Stdout)
		if err != nil {
			utils.PrintError(err)
		}
	} else {
		err := this.RunCmd(os.Args[1], []string{}, os.Stdout)
		if err != nil {
			utils.PrintError(err)
		}
	}
}

func (this *Wel) ServeHTTP(address string) error {
	return this.RunCmd("serve", []string{address}, os.Stdout)
}

func (this *Wel) RunCmd(cmd string, args []string, writer io.Writer) error {
	if lists.ContainsString([]string{"-h", "help", "--help", "?"}, cmd) {

		utils.Println(`wel usage: 
   your-wel [OPTIONS]

OPTIONS:`)
		for _, cmd := range [][]string{
			{"-h", "current help"},
			{"-v", "wel version"},
			{"info", "wel information"},
			{"options", "wel instance options"},
			{"all", "wel all information, options, templates"},
			{"export", "export wel configuration to YAML file"},
			{"pipe", "open a pipe for fetching values"},
			{"fetch -option1=value1 ...", "fetch values"},
			{"run OPERATION -option1=value1 ...", "run instance operation"},
			{"serve HOST:PORT", "serve a http server"},
		} {
			if len(cmd[0]) < 8 {
				fmt.Printf("%-12s %s\n", cmd[0], cmd[1])
			} else {
				fmt.Printf("\n%s\n   %s\n", cmd[0], cmd[1])
			}
		}
	} else if lists.ContainsString([]string{"-v", "version"}, cmd) {
		writer.Write([]byte(this.Name + " v" + this.Version))
	} else if lists.ContainsString([]string{"info"}, cmd) {
		writer.Write(utils.JSONEncodePretty(map[string]interface{}{
			"id":          this.Id,
			"name":        this.Name,
			"developer":   this.Developer,
			"description": this.Description,
			"version":     this.Version,
			"site":        this.Site,
		}))
	} else if lists.ContainsString([]string{"options"}, cmd) {
		writer.Write(utils.JSONEncodePretty(this.Options))
	} else if lists.ContainsString([]string{"all"}, cmd) {
		writer.Write(utils.JSONEncodePretty(map[string]interface{}{
			"id":                 this.Id,
			"name":               this.Name,
			"developer":          this.Developer,
			"description":        this.Description,
			"version":            this.Version,
			"site":               this.Site,
			"options":            this.Options,
			"apps":               this.Apps,
			"operations":         this.Operations,
			"thresholdTemplates": this.ThresholdTemplates,
			"chartTemplates":     this.ChartTemplates,
			"dashboardTemplates": this.DashboardTemplates,
		}))
	} else if lists.ContainsString([]string{"fetch"}, cmd) { // fetch values
		opts := map[string]string{}
		reg := regexp.MustCompile(`^(?:-*)(\w+)=`)
		for _, arg := range args {
			if reg.MatchString(arg) {
				matches := reg.FindStringSubmatch(arg)
				opts[matches[1]] = arg[len(matches[0]):]
			}
		}

		result, err := this.Fetch(opts)
		if err != nil {
			return err
		}
		writer.Write(utils.JSONEncodePretty(result))
	} else if lists.ContainsString([]string{"run"}, cmd) { // run OPERATION OPTIONS
		if len(args) == 0 {
			return errors.New("'OPERATION' should not be empty")
		}

		opCode := args[0]
		if len(opCode) == 0 {
			return errors.New("'OPERATION' should not be empty")
		}

		op := this.FindOperation(opCode)
		if op == nil {
			return errors.New("can not find operation with code '" + opCode + "'")
		}

		handler := op.Handler()
		if handler == nil {
			return errors.New("operation handler should not be nil")
		}

		opts := map[string]string{}
		if len(args) > 1 {
			reg := regexp.MustCompile(`^(?:-*)(\w+)=`)
			for _, arg := range args[1:] {
				if reg.MatchString(arg) {
					matches := reg.FindStringSubmatch(arg)
					opts[matches[1]] = arg[len(matches[0]):]
				}
			}
		}

		result, err := handler(opts)
		if err != nil {
			return err
		}

		writer.Write([]byte(result))
	} else if lists.ContainsString([]string{"serve"}, cmd) { // serve HOST:PORT OPTIONS ...
		if len(args) == 0 {
			return errors.New("'HOST:PORT' should not be empty")
		}

		address := args[0]
		if len(address) == 0 {
			return errors.New("'HOST:PORT' should not be empty")
		}

		server := NewHTTPServer(this, address)
		return server.Start()
	} else if lists.ContainsString([]string{"export"}, cmd) {
		if len(this.Id) == 0 {
			return errors.New("wel id should not be empty")
		}
		data, err := yaml.Marshal(this)
		if err != nil {
			return err
		}
		exe, err := os.Executable()
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(filepath.Dir(exe)+"/"+this.Id+".yaml", data, 0666)
		if err != nil {
			return err
		}
		writer.Write([]byte("'" + this.Id + ".yaml' generated\n"))
		return nil
	} else if lists.ContainsString([]string{"pipe"}, cmd) {
		reader := os.NewFile(uintptr(3), "parentReader")
		if reader == nil {
			return errors.New("reader should not be nil")
		}

		var writer *os.File
		if !utils.IsTesting() {
			writer = os.NewFile(uintptr(4), "parentWriter")
			if writer == nil {
				return errors.New("writer should not be nil")
			}
		} else {
			// for testing
			writer = os.NewFile(uintptr(7), "parentWriter")
			if writer == nil {
				return errors.New("writer should not be nil")
			}
		}

		decoder := msgpack.NewDecoder(reader)
		encoder := msgpack.NewEncoder(writer)

		for {
			cmd := new(CommandMsg)
			err := decoder.Decode(cmd)
			if err != nil {
				if err != io.EOF && err != io.ErrUnexpectedEOF {
					log.Println("[error]" + err.Error())
				}
				break
			}

			reply := new(ReplyMsg)
			reply.Code = cmd.Code
			reply.Id = cmd.Id

			args := []string{}
			for k, v := range cmd.Args {
				args = append(args, k+"="+v)
			}
			output := bytes.NewBuffer([]byte{})
			err = this.RunCmd(cmd.Code, args, output)
			if err != nil {
				reply.Error = err.Error()
			} else {
				reply.Result = string(output.Bytes())
			}

			err = encoder.Encode(reply)
			if err != nil {
				log.Println("[error]" + err.Error())
			}
		}

		return nil
	} else {
		return errors.New("unknown command, use 'wel -h' to lookup usage")
	}
	return nil
}
