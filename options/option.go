package options

import (
	"html"
	"net/http"
	"strings"
)

type OptionInterface interface {
	BasicOption() *Option
	AsHTML() string
	ApplyRequest(req *http.Request) (value interface{}, skip bool, err error)
}

type Option struct {
	Type         string            `yaml:"type" json:"type"` // option type, such as "bool", "string" ...
	Namespace    string            `yaml:"namespace" json:"namespace"`
	Attrs        map[string]string `yaml:"attrs" json:"attrs"`
	Title        string            `yaml:"title" json:"title"`
	Subtitle     string            `yaml:"subtitle" json:"subtitle"`
	Code         string            `yaml:"code" json:"code"`
	Description  string            `yaml:"description" json:"description"`
	IsRequired   bool              `yaml:"isRequired" json:"isRequired"`
	Value        interface{}       `yaml:"value" json:"value"`
	Javascript   string            `yaml:"javascript" json:"javascript"`
	CSS          string            `yaml:"css" json:"css"`
	ValidateCode string            `yaml:"validateCode" json:"validateCode"`
}

func (this *Option) BasicOption() *Option {
	return this
}

func (this *Option) ComposeAttrs(attrs map[string]string) string {
	composedAttrs := map[string]string{}
	for k, v := range this.Attrs {
		composedAttrs[k] = v
	}
	for k, v := range attrs {
		composedAttrs[k] = v
	}

	list := []string{}
	for name, value := range composedAttrs {
		list = append(list, name+"=\""+html.EscapeString(value)+"\"")
	}
	return strings.Join(list, " ")
}

func (this *Option) Attr(name string, value string) {
	if this.Attrs == nil {
		this.Attrs = map[string]string{}
	}
	this.Attrs[name] = value
}
