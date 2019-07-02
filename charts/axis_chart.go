package charts

type AxisChart struct {
	BasicChart `yaml:",inline" json:",inline"`

	Axis *Axis `yaml:"axis" json:"axis"`
}
