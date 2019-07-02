package charts

type TableCol struct {
	Width  string `yaml:"width" json:"width"`
	Header string `yaml:"header" json:"header"`
}

func NewTableCol() *TableCol {
	return &TableCol{}
}

func NewDefaultTableCol() *TableCol {
	return &TableCol{}
}
