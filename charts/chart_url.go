package charts

type URLChart struct {
	BasicChart
}

func (this *URLChart) Type() string {
	return TypeURL
}
