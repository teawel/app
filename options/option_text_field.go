package options

import (
	"fmt"
	"github.com/teawel/app/types"
	"net/http"
)

type TextField struct {
	Option `yaml:",inline"`

	MaxLength   int    `yaml:"maxLength" json:"maxLength"`
	Placeholder string `yaml:"placeholder" json:"placeholder"`
	Size        int    `yaml:"size" json:"size"`
	RightLabel  string `yaml:"rightLabel" json:"rightLabel"`
}

func NewTextField(title string, code string) *TextField {
	return &TextField{
		Option: Option{
			Type:  "string",
			Title: title,
			Code:  code,
		},
	}
}

func (this *TextField) AsHTML() string {
	attrs := map[string]string{}
	if this.MaxLength > 0 {
		attrs["maxlength"] = fmt.Sprintf("%d", this.MaxLength)
	}
	if len(this.Placeholder) > 0 {
		attrs["placeholder"] = this.Placeholder
	}
	attrs["value"] = types.String(this.Value)
	if this.Size > 0 {
		attrs["size"] = fmt.Sprintf("%d", this.Size)
	}
	attrs["name"] = this.Namespace + "_" + this.Code

	if len(this.RightLabel) == 0 {
		return `<input type="text" ` + this.ComposeAttrs(attrs) + ` />`
	} else {
		return `<div class="ui input right labeled"><input type="text" ` + this.ComposeAttrs(attrs) + `> <label class="ui label">` + this.RightLabel + `</label></div>`
	}
}

func (this *TextField) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Namespace + "_" + this.Code)
	return value, false, nil
}
