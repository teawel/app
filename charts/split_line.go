package charts

// gauge split line
type SplitLine struct {
	Length interface{} `yaml:"length" json:"length"`
	Width  interface{} `yaml:"width" json:"width"`
	Color  interface{} `yaml:"color" json:"color"`
}

func NewSplitLine() *SplitLine {
	return &SplitLine{}
}
