package app

import (
	"github.com/teawel/app/charts"
	"github.com/teawel/app/options"
	"github.com/teawel/app/utils"
	"github.com/vmihailenco/msgpack"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
	"time"
)

func TestWelCommandHelp(t *testing.T) {
	wel := NewWel()
	wel.RunCmd("-h", []string{}, os.Stdout)
}

func TestWelCommandUnknown(t *testing.T) {
	wel := NewWel()
	wel.RunCmd("unknown", []string{}, os.Stdout)
}

func TestWelCommandVersion(t *testing.T) {
	wel := NewWel()
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.RunCmd("-v", []string{}, os.Stdout)
}

func TestWelCommandInfo(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	wel.RunCmd("info", []string{}, os.Stdout)
}

func TestWelCommandOptions(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	{
		opt := options.NewStringOption("Host", "host")
		opt.Description = "please input the MySQL server host"
		wel.AddOption(opt)
	}

	{
		opt := options.NewStringOption("Port", "port")
		opt.Description = "please input the MySQL server port"
		wel.AddOption(opt)
	}
	wel.RunCmd("options", []string{}, os.Stdout)
}

func TestWelCommandAll(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	{
		op := NewOperation()
		op.Code = "start"
		op.OnRun(func(options map[string]string) (result string, err error) {
			t.Log("addr:", options["host"]+":"+options["port"])
			result = "success"
			return
		})
		wel.AddOperation(op)
	}
	wel.RunCmd("all", []string{}, os.Stdout)
}

func TestWelCommandFetch(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	{
		opt := options.NewStringOption("Host", "host")
		opt.Description = "please input the MySQL server host"
		wel.AddOption(opt)
	}

	{
		opt := options.NewStringOption("Port", "port")
		opt.Description = "please input the MySQL server port"
		wel.AddOption(opt)
	}
	wel.OnFetch(func(options map[string]string) (result map[string]string, err error) {
		return map[string]string{
			"host": options["host"],
			"port": options["port"],
		}, nil
	})
	err := wel.RunCmd("fetch", []string{"-host=127.0.0.1", "-port=3306"}, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWelCommandOperation(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	{
		op := NewOperation()
		op.Code = "start"
		op.OnRun(func(options map[string]string) (result string, err error) {
			t.Log("addr:", options["host"]+":"+options["port"])
			result = "success"
			return
		})
		wel.AddOperation(op)
	}
	err := wel.RunCmd("run", []string{"start", "-host=127.0.0.1", "-port=3306"}, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWelCommandServe(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	{
		op := NewOperation()
		op.Code = "start"
		op.OnRun(func(options map[string]string) (result string, err error) {
			t.Log("addr:", options["host"]+":"+options["port"])
			result = "success"
			return
		})
		wel.AddOperation(op)
	}
	{
		opt := options.NewStringOption("Host", "host")
		opt.Description = "please input the MySQL server host"
		wel.AddOption(opt)
	}

	{
		opt := options.NewStringOption("Port", "port")
		opt.Description = "please input the MySQL server port"
		wel.AddOption(opt)
	}
	wel.OnFetch(func(options map[string]string) (result map[string]string, err error) {
		res := map[string]string{
			"hello": "World",
		}

		host, ok := options["host"]
		if ok {
			res["host"] = host
		}

		result = res
		return
	})
	err := wel.RunCmd("serve", []string{":8801"}, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWel_ServeHTTP(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql-test@teaos.cn"
	wel.Version = "1.0"
	wel.Name = "MySQL Wel"
	wel.Description = "MySQL statistics, operations"
	wel.Site = "http://wel.teaos.cn"
	wel.Developer = "TeaOS"
	{
		op := NewOperation()
		op.Code = "start"
		op.Name = "Start"
		op.Description = "Start the server"
		op.OnRun(func(options map[string]string) (result string, err error) {
			utils.Println("addr:", options["host"]+":"+options["port"])
			result = "success"
			return
		})
		wel.AddOperation(op)
	}
	{
		op := NewOperation()
		op.Code = "stop"
		op.Name = "Stop"
		op.Description = "Stop the server"
		op.OnRun(func(options map[string]string) (result string, err error) {
			result = "success"
			return
		})
		wel.AddOperation(op)
	}
	{
		app := NewApp()
		app.Name = "MySQL"
		app.Site = "https://mysql.com"
		app.DocumentSite = "https://dev.mysql.com"
		app.SourceSite = "http://github.com"
		app.DownloadSite = "http://download.mysql.com"
		app.Version = "1.0"
		app.Developer = "Oracle"
		app.Description = "a popular database"
		wel.AddApp(app)
	}
	{
		t := NewThreshold()
		t.Expr = "${higher} > 100"
		wel.AddThreshold(t)
	}
	{
		t := NewThreshold()
		t.Expr = "${lower} < 20"
		wel.AddThreshold(t)
	}

	{
		opt := options.NewStringOption("Host", "host")
		opt.Description = "please input the MySQL server host"
		opt.Placeholder = "127.0.0.1"
		wel.AddOption(opt)
	}

	{
		opt := options.NewStringOption("Port", "port")
		opt.Description = "please input the MySQL server port"
		opt.Placeholder = "3306"
		opt.Subtitle = "3306"
		wel.AddOption(opt)
	}

	{
		canvas := NewChartCanvas("cpu_usage", "CPU Usage", CanvasHalf, CanvasFull)
		chart := charts.NewLineChart()
		{
			line := charts.NewLine()
			line.Query = `result("usage")`
			chart.AddLine(line)
		}
		{
			line := charts.NewLine()
			line.Query = `result("usage3")`
			chart.AddLine(line)
		}
		canvas.SetChart(chart)
		wel.AddChart(canvas)
	}

	{
		canvas := NewChartCanvas("memory_usage", "Memory Usage", CanvasHalf, CanvasFull)
		chart := charts.NewLineChart()
		{
			line := charts.NewLine()
			line.Query = `result("usage2")`
			chart.AddLine(line)
		}
		canvas.SetChart(chart)
		wel.AddChart(canvas)
	}

	{
		dashboard := NewDefaultDashboard("1.0.1")
		wel.AddChartsToDashboard(dashboard, "cpu_usage", "memory_usage")
		wel.AddDashboard(dashboard)
	}

	wel.OnFetch(func(options map[string]string) (result map[string]string, err error) {
		res := map[string]string{
			"hello":  "World",
			"usage":  strconv.Itoa(50 + rand.Int()%50),
			"usage2": strconv.Itoa(rand.Int() % 1000),
			"usage3": strconv.Itoa(rand.Int() % 50),
		}

		host, ok := options["host"]
		if ok {
			res["host"] = host
		}

		result = res
		return
	})
	err := wel.ServeHTTP(":8802")
	if err != nil {
		t.Fatal(err)
	}
}

func TestWel_Pipe(t *testing.T) {
	// fd: 3, 4
	_, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	// fd: 6, 7
	r, _, err := os.Pipe()

	go func() {
		time.Sleep(1 * time.Second)
		encoder := msgpack.NewEncoder(w)
		encoder.Encode(&CommandMsg{
			Id:   "123",
			Code: "fetch",
			Args: map[string]string{
				"host": "127.0.0.1",
			},
		})
	}()

	go func() {
		decoder := msgpack.NewDecoder(r)
		for {
			reply := new(ReplyMsg)
			err := decoder.Decode(reply)
			if err != nil {
				t.Fatal(err)
			}
			log.Println("reply:", reply)
		}
	}()

	wel := NewWel()
	wel.OnFetch(func(options map[string]string) (result map[string]string, err error) {
		return map[string]string{
			"name": "lu",
			"age":  "20",
		}, nil
	})
	wel.Id = "mysql-test@teaos.cn"
	t.Log(wel.RunCmd("pipe", []string{}, os.Stdout))
}

func TestWel_Export(t *testing.T) {
	wel := NewWel()
	wel.Id = "mysql"
	wel.Version = "1.0"
	wel.AddOption(options.NewStringOption("Host", "host"))
	wel.AddOption(options.NewStringOption("Port", "port"))
	err := wel.RunCmd("export", []string{}, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("OK")
}
