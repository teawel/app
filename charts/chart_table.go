package charts

type TableChart struct {
	BasicChart `yaml:",inline" json:",inline"`

	Rows []*TableRow `yaml:"rows" json:"rows"`
	Cols []*TableCol `yaml:"cols" json:"cols"`
}

func NewTableChart() *TableChart {
	return &TableChart{
		Rows: []*TableRow{},
		Cols: []*TableCol{},
	}
}

func (this *TableChart) AddRow(row *TableRow) {
	this.Rows = append(this.Rows, row)
}

func (this *TableChart) AddRowValues(value ...string) {
	this.Rows = append(this.Rows, NewTableRow(value...))
}

func (this *TableChart) AddCol(col *TableCol) {
	this.Cols = append(this.Cols, col)
}

func (this *TableChart) AddDefaultCol() {
	this.Cols = append(this.Cols, NewDefaultTableCol())
}

func (this *TableChart) Type() string {
	return TypeTable
}
