package charts

import "encoding/json"

type Type = string

const (
	TypeValue    = "value"
	TypeLine     = "line"
	TypePie      = "pie"
	TypeBar      = "bar"
	TypeStackBar = "stackBar"
	TypeFunnel   = "funnel"
	TypeRadar    = "radar"
	TypeScatter  = "scatter"
	TypeGauge    = "gauge"
	TypeEchart   = "echart"
	TypeTable    = "table"
	TypeClock    = "clock"
	TypeURL      = "url"
	TypeHTML     = "html"
)

func Decode(chartType Type, options string) (ChartInterface, error) {
	var chart ChartInterface = nil
	switch chartType {
	case TypeValue:
		chart = new(ValueChart)
	case TypeLine:
		chart = new(LineChart)
	case TypePie:
		chart = new(PieChart)
	case TypeBar:
		chart = new(BarChart)
	case TypeStackBar:
		chart = new(StackBarChart)
	case TypeFunnel:
		chart = new(FunnelChart)
	case TypeRadar:
		chart = new(RadarChart)
	case TypeScatter:
		chart = new(ScatterChart)
	case TypeGauge:
		chart = new(GaugeChart)
	case TypeEchart:
		chart = new(Echart)
	case TypeTable:
		chart = new(TableChart)
	case TypeClock:
		chart = new(ClockChart)
	case TypeURL:
		chart = new(URLChart)
	case TypeHTML:
		chart = new(HTMLChart)
	}
	err := json.Unmarshal([]byte(options), chart)
	if err != nil {
		return nil, err
	}
	return chart, nil
}
