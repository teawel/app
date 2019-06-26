package charts

type Echart struct {
	BasicChart
}

func (this *Echart) Type() string {
	return TypeEchart
}
