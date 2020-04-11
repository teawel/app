package options

import (
	"fmt"
	"github.com/teawel/app/types"
	"net/http"
)

type TextBox struct {
	Option

	MaxLength   int    `yaml:"maxLength" json:"maxLength"`
	Placeholder string `yaml:"placeholder" json:"placeholder"`
	Cols        int    `yaml:"cols" json:"cols"`
	Rows        int    `yaml:"rows" json:"rows"`
}

func NewTextBox(title string, code string) *TextBox {
	return &TextBox{
		Option: Option{
			Type:  "memo",
			Title: title,
			Code:  code,
		},
	}
}

func (this *TextBox) AsHTML() string {
	attrs := map[string]string{}
	if this.MaxLength > 0 {
		attrs["maxlength"] = fmt.Sprintf("%d", this.MaxLength)
	}

	if len(this.Placeholder) > 0 {
		attrs["placeholder"] = this.Placeholder
	}

	if this.Cols > 0 {
		attrs["cols"] = fmt.Sprintf("%d", this.Cols)
	}

	if this.Rows > 0 {
		attrs["rows"] = fmt.Sprintf("%d", this.Rows)
	}

	valueString := types.String(this.Value)
	attrs["name"] = this.Namespace + "_" + this.Code

	return `<textarea ` + this.ComposeAttrs(attrs) + `>` + valueString + `</textarea>`
}

func (this *TextBox) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Namespace + "_" + this.Code)
	return value, false, nil
}
