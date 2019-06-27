package charts

type ClockChart struct {
	BasicChart

	Timestamp int64 `yaml:"timestamp" json:"timestamp"`
}

func (this *ClockChart) Type() string {
	return TypeClock
}
