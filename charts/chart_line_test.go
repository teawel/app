package charts

import (
	"github.com/teawel/app/utils"
	"testing"
)

func TestNewLineChart(t *testing.T) {
	chart := NewLineChart()
	chart.MaxValue = 1024

	{
		line := NewLine()
		line.AddValue(1, "")
		line.AddValue(2, "")
		line.AddValue(3, "end")
		chart.AddLine(line)
	}

	t.Log(string(utils.JSONEncodePretty(chart)))
}
