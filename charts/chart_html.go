package charts

type HTMLChart struct {
	BasicChart
}

func (this *HTMLChart) Type() string {
	return TypeHTML
}
