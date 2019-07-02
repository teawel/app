package charts

type XAxis struct {
	Name   string      `yaml:"name" json:"name"`
	Type   string      `yaml:"type" json:"type"`
	Max    interface{} `yaml:"max" json:"max"`
	Labels []string    `yaml:"labels" json:"labels"`
}

func NewXAxis() *XAxis {
	return &XAxis{}
}
