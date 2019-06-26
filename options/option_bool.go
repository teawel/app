package options

import (
	"github.com/teawel/app/types"
	"net/http"
)

type BoolOption struct {
	Option

	IsChecked bool   `yaml:"isChecked" json:"isChecked"`
	Label     string `yaml:"label" json:"label"`
}

func NewBoolOption(title string, code string) *BoolOption {
	return &BoolOption{
		Option: Option{
			Type:  "bool",
			Title: title,
			Code:  code,
		},
	}
}

func (this *BoolOption) AsHTML() string {
	attrs := map[string]string{
		"name": this.Namespace + "_" + this.Code,
	}
	if this.IsChecked {
		attrs["checked"] = "checked"
	}
	attrs["value"] = types.String(this.Value)
	return `
<div class="ui checkbox">
<input type="checkbox"` + this.ComposeAttrs(attrs) + `/>
<label>` + this.Label + `</label>
</div>
`
}

func (this *BoolOption) ApplyRequest(req *http.Request) (value interface{}, skip bool, err error) {
	value = req.Form.Get(this.Namespace + "_" + this.Code)
	return value, false, nil
}
