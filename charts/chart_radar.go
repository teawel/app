package charts

type RadarChart struct {
	BasicChart

	Radius     string           `yaml:"radius" json:"radius"`
	Center     *Position        `yaml:"center" json:"center"`
	Categories []*RadarCategory `yaml:"categories" json:"categories"`
	Series     []*RadarSeries   `yaml:"series" json:"series"`
}

func NewRadarChart() *RadarChart {
	return &RadarChart{
		Categories: []*RadarCategory{},
		Series:     []*RadarSeries{},
	}
}

func (this *RadarChart) AddCategory(category *RadarCategory) {
	this.Categories = append(this.Categories, category)
}

func (this *RadarChart) AddSeries(series *RadarSeries) {
	this.Series = append(this.Series, series)
}

func (this *RadarChart) Type() string {
	return TypeRadar
}
