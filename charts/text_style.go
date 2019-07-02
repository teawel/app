package charts

type TextStyle struct {
	Color      interface{} `yaml:"color" json:"color"`
	FontSize   interface{} `yaml:"fontSize" json:"fontSize"`
	FontFamily interface{} `yaml:"fontFamily" json:"fontFamily"`
	FontWeight interface{} `yaml:"fontWeight" json:"fontWeight"`
}

func NewTextStyle() *TextStyle {
	return &TextStyle{}
}
