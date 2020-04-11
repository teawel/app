package options

import (
	"net/http"
)

type OptionInterface interface {
	BasicOption() *Option
	ApplyRequest(req *http.Request) (value interface{}, skip bool, err error)
}

type Option struct {
	Type        string            `yaml:"type" json:"type"` // option type, such as "bool", "string" ...
	Attrs       map[string]string `yaml:"attrs" json:"attrs"`
	Title       string            `yaml:"title" json:"title"`
	Subtitle    string            `yaml:"subtitle" json:"subtitle"`
	Code        string            `yaml:"code" json:"code"`
	Description string            `yaml:"description" json:"description"`
	IsRequired  bool              `yaml:"isRequired" json:"isRequired"`
	Value       interface{}       `yaml:"value" json:"value"`
}

func (this *Option) BasicOption() *Option {
	return this
}

func (this *Option) Attr(name string, value string) {
	if this.Attrs == nil {
		this.Attrs = map[string]string{}
	}
	this.Attrs[name] = value
}
