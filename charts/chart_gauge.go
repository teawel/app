package charts

type GaugeChart struct {
	BasicChart
}

func (this *GaugeChart) Type() string {
	return TypeGauge
}
