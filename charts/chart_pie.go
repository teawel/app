package charts

type PieChart struct {
	BasicChart

	Values []*Value `yaml:"values" json:"values"`
	Colors []string `yaml:"colors" json:"colors"`
}

func (this *PieChart) Type() string {
	return TypePie
}
