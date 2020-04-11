package options

import (
	"net/http"
)

type Checkbox struct {
	Option

	IsChecked bool   `yaml:"isChecked" json:"isChecked"`
	Label     string `yaml:"label" json:"label"`
}

func NewCheckbox(title string, code string) *Checkbox {
	return &Checkbox{
		Option: Option{
			Type:  "bool",
			Title: title,
			Code:  code,
		},
	}
}

func (this *Checkbox) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Code)
	return value, false, nil
}
