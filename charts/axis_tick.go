package charts

type AxisTick struct {
	Length interface{} `yaml:"length" json:"length"`
	Width  interface{} `yaml:"width" json:"width"`
	Color  interface{} `yaml:"color" json:"color"`
}

func NewAxisTick() *AxisTick {
	return &AxisTick{}
}
