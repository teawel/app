package charts

type Axis struct {
	Reverse bool   `yaml:"reverse" json:"reverse"`
	X       *XAxis `yaml:"x" json:"x"`
	Y       *YAxis `yaml:"y" json:"y"`
}

func NewAxis() *Axis {
	return &Axis{
		X: NewXAxis(),
		Y: NewYAxis(),
	}
}
