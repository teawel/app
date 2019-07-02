package charts

// gauge pointer
type Pointer struct {
	Length interface{} `yaml:"length" json:"length"`
	Width  interface{} `yaml:"width" json:"width"`
}

func NewPointer() *Pointer {
	return &Pointer{}
}
