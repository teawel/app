package charts

type BarChart struct {
	AxisChart `yaml:",inline" json:",inline"`

	Labels []string     `yaml:"labels" json:"labels"`
	Series []*BarSeries `yaml:"series" json:"series"`
}

func NewBarChart() *BarChart {
	return &BarChart{
		Labels: []string{},
		Series: []*BarSeries{},
	}
}

func (this *BarChart) AddSeries(series *BarSeries) {
	this.Series = append(this.Series, series)
}

func (this *BarChart) Type() string {
	return TypeBar
}
