package charts

type BarChart struct {
	BasicChart
}

func (this *BarChart) Type() string {
	return TypeBar
}
