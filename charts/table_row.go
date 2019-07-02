package charts

type TableRow struct {
	Values []string `yaml:"values" json:"values"`
}

func NewTableRow(value ...string) *TableRow {
	return &TableRow{
		Values: value,
	}
}
