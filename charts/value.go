package charts

type Value struct {
	Value interface{} `yaml:"value" json:"value"` // raw value
	Label string      `yaml:"label" json:"label"`
}

func NewValue(value interface{}, label string) *Value {
	return &Value{
		Value: value,
		Label: label,
	}
}
