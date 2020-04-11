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
		opt := options.NewTextField("Host", "host")
		opt.Description = "please input the MySQL server host"
		wel.AddOption(opt)
	}

	{
		opt := options.NewTextField("Port", "port")
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
		opt := options.NewTextField("Host", "host")
		opt.Description = "please input the MySQL server host"
		wel.AddOption(opt)
	}

	{
		opt := options.NewTextField("Port", "port")
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
		opt := options.NewTextField("Host", "host")
		opt.Description = "please input the MySQL server host"
		wel.AddOption(opt)
	}

	{
		opt := options.NewTextField("Port", "port")
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
		opt := options.NewTextField("Host", "host")
		opt.Description = "please input the MySQL server host"
		opt.Placeholder = "127.0.0.1"
		wel.AddOption(opt)
	}

	{
		opt := options.NewTextField("Port", "port")
		opt.Description = "please input the MySQL server port"
		opt.Placeholder = "3306"
		opt.Subtitle = "3306"
		wel.AddOption(opt)
	}

	{
		canvas := NewChartCanvas("cpu_usage", "CPU Usage", CanvasHalf, CanvasFull)
		chart := charts.NewLineChart()
		{
			line := charts.NewLineSeries()
			line.Query = `result("usage")`
			chart.AddSeries(line)
		}
		{
			line := charts.NewLineSeries()
			line.Query = `result("usage3")`
			chart.AddSeries(line)
		}
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	{
		canvas := NewChartCanvas("memory_usage", "Memory Usage", CanvasHalf, CanvasFull)
		chart := charts.NewLineChart()
		{
			line := charts.NewLineSeries()
			line.Query = `result("usage2")`
			chart.AddSeries(line)
		}
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	{
		canvas := NewChartCanvas("single_value", "Single Value", CanvasHalf, CanvasFull)
		chart := charts.NewValueChart()
		//chart.Value = charts.NewValue(1024, "LABEL")
		chart.ValueColor = "red"
		chart.LabelColor = "blue"
		chart.Query = `label(function () {
	return "LABEL";
}).result("usage")`
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	/**{
		canvas := NewChartCanvas("chart_url", "URL Chart", CanvasHalf, CanvasFull)
		chart := charts.NewURLChart()
		chart.URL = "http://teaos.cn"
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}**/

	{
		canvas := NewChartCanvas("chart_html", "HTML Chart", CanvasHalf, CanvasFull)
		chart := charts.NewHTMLChart()
		chart.HTML = "<strong>THIS IS <span style=\"color:red\">H</span><span style=\"color:blue\">T</span>ML</strong>"
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// pie
	{
		canvas := NewChartCanvas("chart_pie", "Pie Chart", CanvasHalf, CanvasFull)
		chart := charts.NewPieChart()
		{
			series := charts.NewPieSeries()
			series.Query = `label(function (k, v) {
	return ["Usage", "Usage2", "Usage3"];
}).result(["usage", "usage2", "usage3"])`
			chart.AddSeries(series)
		}
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// clock
	{
		canvas := NewChartCanvas("chart_clock", "Clock Chart", CanvasHalf, CanvasFull)
		chart := charts.NewClockChart()
		chart.Timestamp = time.Now().Unix()
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// radar
	{
		canvas := NewChartCanvas("chart_radar", "Radar Chart", CanvasHalf, CanvasFull)
		chart := charts.NewRadarChart()
		chart.AddCategory(charts.NewRadarCategory("Sales", 10000))
		chart.AddCategory(charts.NewRadarCategory("Administration", 10000))
		chart.AddCategory(charts.NewRadarCategory("Information", 10000))
		chart.AddCategory(charts.NewRadarCategory("Customer", 10000))
		chart.AddCategory(charts.NewRadarCategory("Marketing", 10000))

		series := charts.NewRadarSeries()
		{
			subCategory := charts.NewRadarSubCategory("Budget")
			subCategory.Values = []float64{4000, 5000, 6000, 7000, 6500}
			series.AddSubCategory(subCategory)
		}

		{
			subCategory := charts.NewRadarSubCategory("Spending")
			subCategory.Values = []float64{3500, 4600, 5000, 3000, 6000}
			series.AddSubCategory(subCategory)
		}

		chart.AddSeries(series)
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// funnel
	{
		canvas := NewChartCanvas("chart_funnel", "Funnel Chart", CanvasHalf, CanvasFull)
		chart := charts.NewFunnelChart()
		series := charts.NewFunnelSeries()
		series.AddValue(60, "Visiting")
		series.AddValue(40, "Consulting")
		series.AddValue(20, "Order")
		series.AddValue(80, "Clicking")
		series.AddValue(100, "Display")
		chart.AddSeries(series)
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// scatter
	{
		canvas := NewChartCanvas("chart_scatter", "Scatter Chart", CanvasHalf, CanvasFull)
		chart := charts.NewScatterChart()

		{
			series := charts.NewScatterSeries()
			series.AddValue(60, "Visiting")
			series.AddValue(40, "Consulting")
			series.AddValue(20, "Order")
			series.AddValue(80, "Clicking")
			series.AddValue(100, "Display")
			series.Size = "20"
			chart.AddSeries(series)
		}

		{
			series := charts.NewScatterSeries()
			series.AddValue(50, "Visiting")
			series.AddValue(45, "Consulting")
			series.AddValue(23, "Order")
			series.AddValue(8, "Clicking")
			series.AddValue(53, "Display")
			series.Size = charts.NewFuncString(`function (value) {
	return value.value / 3;
}`)
			chart.AddSeries(series)
		}

		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// bar
	{
		canvas := NewChartCanvas("chart_bar", "Bar Chart", CanvasHalf, CanvasFull)
		chart := charts.NewBarChart()
		{
			series := charts.NewBarSeries()
			series.AddValue("10", "PHP")
			series.AddValue("20", "Java")
			series.AddValue("30", "Python")
			series.AddValue("25", "Golang")
			series.AddValue("15", "Perl")
			series.Width = "20"
			chart.AddSeries(series)
		}
		{
			series := charts.NewBarSeries()
			series.AddValue("18", "PHP")
			series.AddValue("22", "Java")
			series.AddValue("38", "Python")
			series.AddValue("20", "Golang")
			series.AddValue("19", "Perl")
			chart.AddSeries(series)
		}
		chart.Labels = []string{"PHP", "Java", "Python", "Golang", "Perl"}
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// stack bar
	{
		canvas := NewChartCanvas("chart_stack_bar", "Stack Bar Chart", CanvasHalf, CanvasFull)
		chart := charts.NewBarChart()

		axis := charts.NewAxis()
		axis.Reverse = true
		axis.X.Max = 100
		axis.X.Name = "Usage"
		axis.Y.Name = "Disk"
		chart.Axis = axis

		{
			series := charts.NewBarSeries()
			series.Name = "Size"
			series.AddValue("10", "C:")
			series.AddValue("20", "D:")
			series.AddValue("30", "E:")
			series.Stack = "usage"
			chart.AddSeries(series)
		}
		{
			series := charts.NewBarSeries()
			series.Name = "Used"
			series.AddValue("8", "C:")
			series.AddValue("18", "D:")
			series.AddValue("15", "E:")
			series.Stack = "usage"
			chart.AddSeries(series)
		}
		chart.Labels = []string{"C:", "D:", "E:"}
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// gauge
	{
		canvas := NewChartCanvas("chart_gauge", "Gauge", CanvasHalf, CanvasFull)
		chart := charts.NewGaugeChart()
		{
			series := charts.NewGaugeSeries()
			series.AddValue(100, "km/h")
			series.Min = 10
			series.Max = 200
			series.Radius = "70%"
			series.Center = charts.NewPosition("50%", "50%")

			axisLine := charts.NewAxisLine()
			axisLine.Width = 5
			series.AxisLine = axisLine

			axisTick := charts.NewAxisTick()
			axisTick.Length = 10
			//axisTick.Color = "red"
			//axisTick.Width = 10
			series.AxisTick = axisTick

			splitLine := charts.NewSplitLine()
			splitLine.Length = 10
			series.SplitLine = splitLine

			series.SplitNumber = 5

			pointer := charts.NewPointer()
			pointer.Length = "80%"
			pointer.Width = 2
			series.Pointer = pointer

			series.Detail = charts.NewFuncString(`
function (value) {
	return value + '(Speed)';
}
`)

			detailStyle := charts.NewTextStyle()
			detailStyle.FontSize = 10
			series.DetailStyle = detailStyle

			chart.AddSeries(series)
		}
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	// table
	{
		canvas := NewChartCanvas("chart_table", "Table", CanvasHalf, CanvasFull)

		table := charts.NewTableChart()
		table.AddDefaultCol()
		table.AddDefaultCol()

		{
			col := charts.NewTableCol()
			col.Header = "Location"
			col.Width = "10em"
			table.AddCol(col)
		}

		table.AddRowValues("lu", "20", "Beijing")
		table.AddRowValues("ping", "21", "Shanghai")
		table.AddRowValues("wei", "22", "Shenzhen")
		table.AddRowValues("yin", "21", "Nanjing")
		table.AddRowValues("feng", "25", "Wuxi")

		canvas.SetChart(table)
		wel.AddChartCanvas(canvas)
	}

	// echart
	{
		canvas := NewChartCanvas("chart_echart", "EChart", CanvasHalf, CanvasFull)
		chart := charts.NewEchart()
		chart.Code = `var data = [];

for (var i = 0; i <= 360; i++) {
    var t = i / 180 * Math.PI;
    var r = Math.sin(2 * t) * Math.cos(2 * t);
    data.push([r, i]);
}

var option = {
    title: {
        text: ''
    },
    legend: {
        data: ['line']
    },
    polar: {
        center: ['50%', '54%']
    },
    tooltip: {
        trigger: 'axis',
        axisPointer: {
            type: 'cross'
        }
    },
    angleAxis: {
        type: 'value',
        startAngle: 0
    },
    radiusAxis: {
        min: 0
    },
    series: [{
        coordinateSystem: 'polar',
        name: 'line',
        type: 'line',
        showSymbol: false,
        data: data
    }],
    animationDuration: 2000
};
chart.setOption(option);
`
		canvas.SetChart(chart)
		wel.AddChartCanvas(canvas)
	}

	{
		dashboard := NewDefaultDashboard("1.0.25")
		wel.AddAllChartsToDashboard(dashboard)
		wel.AddDashboard(dashboard)
	}

	{
		dashboard := NewDashboard("menu", "1.0", "Menu")

		{
			canvas := NewChartCanvas("single_value", "Single Value", CanvasFull, CanvasFull)

			{
				menu := charts.NewMenu()
				menu.AddItem(charts.NewMenuItem("kb", "KB"))
				menu.AddItem(charts.NewMenuItem("mb", "MB"))
				menu.AddItem(charts.NewMenuItem("gb", "GB"))
				menu.SelectItem("kb")
				canvas.LeftMenu = menu
			}

			{
				menu := charts.NewMenu()
				menu.AddItem(charts.NewMenuItem("hourly", "Hourly"))
				menu.AddItem(charts.NewMenuItem("daily", "Daily"))
				menu.AddItem(charts.NewMenuItem("weekly", "Weekly"))
				menu.AddItem(charts.NewMenuItem("monthly", "Monthly"))
				menu.SelectItem("hourly")
				canvas.RightMenu = menu
			}

			chart := charts.NewValueChart()
			chart.Value = charts.NewValue("1024", "KB")
			canvas.SetChart(chart)

			dashboard.AddChart(canvas)
		}

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
	wel.AddOption(options.NewTextField("Host", "host"))
	wel.AddOption(options.NewTextField("Port", "port"))
	err := wel.RunCmd("export", []string{}, os.Stdout)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("OK")
}
