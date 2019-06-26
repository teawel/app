package charts

type ScatterChart struct {
	BasicChart
}

func (this *ScatterChart) Type() string {
	return TypeScatter
}
