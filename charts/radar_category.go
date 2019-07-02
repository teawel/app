package charts

type RadarCategory struct {
	Name string  `yaml:"name" json:"name"`
	Max  float64 `yaml:"max" json:"max"`
}

func NewRadarCategory(name string, max float64) *RadarCategory {
	return &RadarCategory{
		Name: name,
		Max:  max,
	}
}
