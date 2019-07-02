package charts

type URLChart struct {
	BasicChart `yaml:",inline" json:",inline"`

	URL string `yaml:"url" json:"url"`
}

func NewURLChart() *URLChart {
	return &URLChart{}
}

func (this *URLChart) Type() string {
	return TypeURL
}
