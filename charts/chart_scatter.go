package charts

type ScatterChart struct {
	AxisChart `yaml:",inline" json:",inline"`

	Series []*ScatterSeries `yaml:"series" json:"series"`
}

func NewScatterChart() *ScatterChart {
	return &ScatterChart{
		Series: []*ScatterSeries{},
	}
}

func (this *ScatterChart) AddSeries(series *ScatterSeries) {
	this.Series = append(this.Series, series)
}

func (this *ScatterChart) Type() string {
	return TypeScatter
}
