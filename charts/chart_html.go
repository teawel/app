package charts

type HTMLChart struct {
	BasicChart

	HTML string `yaml:"html" json:"html"`
}

func NewHTMLChart() *HTMLChart {
	return &HTMLChart{}
}

func (this *HTMLChart) Type() string {
	return TypeHTML
}
