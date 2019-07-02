package charts

type LineSeries struct {
	Series

	IsFilled bool `yaml:"isFilled" json:"isFilled"`
	IsSmooth bool `yaml:"isSmooth" json:"isSmooth"`
}

func NewLineSeries() *LineSeries {
	return &LineSeries{
		Series: Series{
			Values: []*Value{},
		},
	}
}
