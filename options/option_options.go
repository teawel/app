package options

import (
	"github.com/teawel/app/types"
	"html"
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

func (this *Options) AsHTML() string {
	this.Attr("name", this.Namespace+"_"+this.Code)
	result := `<select ` + this.ComposeAttrs(this.Attrs) + ` class="ui dropdown">` + "\n"
	for _, o := range this.Values {
		result += `<option value="` + html.EscapeString(o.Value) + `"`
		if o.Value == types.String(this.Value) {
			result += ` selected="selected"`
		}
		result += ">" + html.EscapeString(o.Text) + `</option>` + "\n"
	}
	result += "</select>"
	return result
}

func (this *Options) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Namespace + "_" + this.Code)
	return value, false, nil
}
