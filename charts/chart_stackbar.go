package charts

type StackBarChart struct {
	BasicChart
}

func (this *StackBarChart) Type() string {
	return TypeStackBar
}
