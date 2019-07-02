package charts

type RadarSeries struct {
	Series

	SubCategories []*RadarSubCategory `yaml:"subCategories" json:"subCategories"`
}

func NewRadarSeries() *RadarSeries {
	return &RadarSeries{
		SubCategories: []*RadarSubCategory{},
	}
}

func (this *RadarSeries) AddSubCategory(category *RadarSubCategory) {
	this.SubCategories = append(this.SubCategories, category)
}
