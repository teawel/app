package options

import (
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

func (this *TextBox) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Code)
	return value, false, nil
}
