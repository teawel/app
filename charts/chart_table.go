package charts

type TableChart struct {
	BasicChart
}

func (this *TableChart) Type() string {
	return TypeTable
}
