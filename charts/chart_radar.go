package charts

type RadarChart struct {
	BasicChart
}

func (this *RadarChart) Type() string {
	return TypeRadar
}
