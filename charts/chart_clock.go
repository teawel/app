package charts

type ClockChart struct {
	BasicChart
}

func (this *ClockChart) Type() string {
	return TypeClock
}
