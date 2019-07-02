package charts

type RadarSubCategory struct {
	Name   string    `yaml:"name" json:"name"`
	Values []float64 `yaml:"values" json:"values"`
}

func NewRadarSubCategory(name string) *RadarSubCategory {
	return &RadarSubCategory{
		Name:   name,
		Values: []float64{},
	}
}
