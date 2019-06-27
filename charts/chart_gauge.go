package charts

type GaugeChart struct {
	BasicChart

	Value float64 `yaml:"value" json:"value"`
	Label string  `yaml:"label" json:"label"`
	Min   float64 `yaml:"min" json:"min"`
	Max   float64 `yaml:"max" json:"max"`
	Unit  string  `yaml:"unit" json:"unit"`
}

func (this *GaugeChart) Type() string {
	return TypeGauge
}
