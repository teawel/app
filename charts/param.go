package charts

type Param struct {
	MenuParam *MenuParam `yaml:"menuParam" json:"menuParam"`
	TimeParam *TimeParam `yaml:"timeParam" json:"timeParam"`
}

func NewParam() *Param {
	return &Param{}
}
