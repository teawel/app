package charts

type ValueChart struct {
	BasicChart

	Value *Value
}

func (this *ValueChart) Type() string {
	return TypeValue
}
