package charts

type ClockChart struct {
	BasicChart

	Timestamp int64 `yaml:"timestamp" json:"timestamp"`
}

func NewClockChart() *ClockChart {
	return &ClockChart{}
}

func (this *ClockChart) Type() string {
	return TypeClock
}
