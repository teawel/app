package charts

type PieChart struct {
	BasicChart
}

func (this *PieChart) Type() string {
	return TypePie
}
