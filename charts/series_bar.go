package charts

type BarSeries struct {
	Series

	Width string `yaml:"width" json:"width"` // bar width
	Stack string `yaml:"stack" json:"stack"`
}

func NewBarSeries() *BarSeries {
	return &BarSeries{
		Series: Series{
			Values: []*Value{},
		},
	}
}
