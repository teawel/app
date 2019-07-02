package charts

type FunnelChart struct {
	BasicChart

	Series []*FunnelSeries `yaml:"series" json:"series"`
}

func NewFunnelChart() *FunnelChart {
	return &FunnelChart{}
}

func (this *FunnelChart) AddSeries(series *FunnelSeries) {
	this.Series = append(this.Series, series)
}

func (this *FunnelChart) Type() string {
	return TypeFunnel
}
