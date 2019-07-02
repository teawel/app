package charts

type ScatterSeries struct {
	Series

	Size string `yaml:"size" json:"size"`
}

func NewScatterSeries() *ScatterSeries {
	return &ScatterSeries{
		Series: Series{
			Values: []*Value{},
		},
	}
}
