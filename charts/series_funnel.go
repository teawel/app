package charts

type FunnelSeries struct {
	Series
}

func NewFunnelSeries() *FunnelSeries {
	return &FunnelSeries{
		Series{
			Values: []*Value{},
		},
	}
}
