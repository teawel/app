package app

type Threshold struct {
	Expr    string          `yaml:"expr" json:"expr"`
	Level   ThresholdLevel  `yaml:"level" json:"level"`
	Actions []*ActionConfig `yaml:"actions" json:"actions"`
}

func NewThreshold() *Threshold {
	return &Threshold{
		Level:   ThresholdLevelInfo,
		Actions: []*ActionConfig{},
	}
}
