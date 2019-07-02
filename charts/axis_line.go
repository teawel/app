package charts

type AxisLine struct {
	Width interface{} `yaml:"width" json:"width"`
}

func NewAxisLine() *AxisLine {
	return &AxisLine{}
}
