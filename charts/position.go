package charts

type Position struct {
	X interface{} `yaml:"x" json:"x"`
	Y interface{} `yaml:"y" json:"y"`
}

func NewPosition(x, y interface{}) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}
