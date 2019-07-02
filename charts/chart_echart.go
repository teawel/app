package charts

type Echart struct {
	BasicChart

	Code string `yaml:"code" json:"code"`
}

func NewEchart() *Echart {
	return &Echart{}
}

func (this *Echart) Type() string {
	return TypeEchart
}
