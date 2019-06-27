package charts

type URLChart struct {
	BasicChart

	URL string `yaml:"url" json:"url"`
}

func (this *URLChart) Type() string {
	return TypeURL
}
