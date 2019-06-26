package charts

type FunnelChart struct {
	BasicChart
}

func (this *FunnelChart) Type() string {
	return TypeFunnel
}
