package app

import (
	"github.com/teawel/app/charts"
	"github.com/teawel/app/dbs"
	"github.com/teawel/app/utils"
	"os"
	"testing"
	"time"
)

func TestNewChart(t *testing.T) {
	canvas := NewChartCanvas("chart1", "cpu usage", CanvasHalf, CanvasFull)

	chart := charts.NewLineChart()
	{
		line := charts.NewLine()
		chart.AddLine(line)
	}
	canvas.SetChart(chart)

	t.Log(string(utils.JSONEncodePretty(canvas)))
	t.Log(canvas.OptionsJSON)
}

func TestChart_Values(t *testing.T) {
	before := time.Now()

	db, err := dbs.NewDB(os.TempDir() + "test.db")
	if err != nil {
		t.Fatal(err)
	}

	chart := charts.NewLineChart()
	{
		line := charts.NewLine()
		line.Query = `
	filter(function (k, v) {
		return true;
	}).
	map(function (k, v) {
		return v;
	}).
	result("name")
`
		chart.AddLine(line)
	}

	err = chart.Fetch(db)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(utils.JSONEncodePretty(chart)))
	t.Logf("%.6fms", time.Since(before).Seconds()*100)
}
