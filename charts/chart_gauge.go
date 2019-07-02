package charts

type GaugeChart struct {
	BasicChart

	Series []*GaugeSeries `yaml:"series" json:"series"`
}

func NewGaugeChart() *GaugeChart {
	return &GaugeChart{
		Series: []*GaugeSeries{},
	}
}

func (this *GaugeChart) AddSeries(series *GaugeSeries) {
	this.Series = append(this.Series, series)
}

func (this *GaugeChart) Type() string {
	return TypeGauge
}
