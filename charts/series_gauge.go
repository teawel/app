package charts

type GaugeSeries struct {
	Series `yaml:",inline" json:",inline"`

	Min         interface{} `yaml:"min" json:"min"`
	Max         interface{} `yaml:"max" json:"max"`
	Radius      interface{} `yaml:"radius" json:"radius"`
	Center      *Position   `yaml:"center" json:"center"`
	AxisLine    *AxisLine   `yaml:"axisLine" json:"axisLine"`
	AxisTick    *AxisTick   `yaml:"axisTick" json:"axisTick"`
	Pointer     *Pointer    `yaml:"pointer" json:"pointer"`
	SplitLine   *SplitLine  `yaml:"splitLine" json:"splitLine"`
	SplitNumber interface{} `yaml:"splitNumber" json:"splitNumber"`
	Detail      interface{} `yaml:"detail" json:"detail"`
	DetailStyle *TextStyle  `yaml:"detailStyle" json:"detailStyle"`
}

func NewGaugeSeries() *GaugeSeries {
	return &GaugeSeries{}
}
