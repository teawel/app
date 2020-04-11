package options

import (
	"net/http"
)

type Options struct {
	Option

	Values []*TextValue `yaml:"values" json:"values"`
}

type TextValue struct {
	Text  string `yaml:"text" json:"text"`
	Value string `yaml:"value" json:"value"`
}

func NewOptions(title string, code string) *Options {
	return &Options{
		Option: Option{
			Type:  "options",
			Title: title,
			Code:  code,
		},
	}
}

func (this *Options) AddValue(text string, value string) {
	this.Values = append(this.Values, &TextValue{
		Text:  text,
		Value: value,
	})
}

func (this *Options) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Code)
	return value, false, nil
}
