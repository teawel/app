package charts

type Line struct {
	Name     string   `yaml:"name" json:"name"`
	Values   []*Value `yaml:"values" json:"values"`
	IsFilled bool     `yaml:"isFilled" json:"isFilled"`
	IsSmooth bool     `yaml:"isSmooth" json:"isSmooth"`
	Color    string   `yaml:"color" json:"color"`
	Query    string   `yaml:"query" json:"query"`
}

func NewLine() *Line {
	return &Line{
		Values: []*Value{},
	}
}

func (this *Line) AddValues(values []*Value) {
	this.Values = append(this.Values, values...)
}

func (this *Line) AddValue(value interface{}, label string) {
	this.Values = append(this.Values, NewValue(value, label))
}
