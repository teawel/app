package charts

type Echart struct {
	BasicChart

	Code string `yaml:"code" json:"code"`
}

func (this *Echart) Type() string {
	return TypeEchart
}
