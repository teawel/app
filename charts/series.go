package charts

type Series struct {
	Name   string   `yaml:"name" json:"name"`
	Values []*Value `yaml:"values" json:"values"`
	Color  string   `yaml:"color" json:"color"`
	Query  string   `yaml:"query" json:"query"`
	//Z      int      `yaml:"z" json:"z"`
}

func (this *Series) AddValues(values []*Value) {
	this.Values = append(this.Values, values...)
}

func (this *Series) AddValue(value interface{}, label string) {
	this.Values = append(this.Values, NewValue(value, label))
}
