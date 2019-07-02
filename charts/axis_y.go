package charts

type YAxis struct {
	XAxis `yaml:",inline" json:",inline"`
}

func NewYAxis() *YAxis {
	return &YAxis{}
}
