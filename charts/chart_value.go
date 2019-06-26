package charts

type ValueChart struct {
	BasicChart
}

func (this *ValueChart) Type() string {
	return TypeValue
}
