package charts

type TableChart struct {
	BasicChart

	Rows      [][]string `yaml:"rows" json:"rows"`
	ColWidths []float64  `yaml:"colWidths" json:"colWidths"`
}

func (this *TableChart) Type() string {
	return TypeTable
}
