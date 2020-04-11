package options

import (
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

func (this *TextField) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Code)
	return value, false, nil
}
