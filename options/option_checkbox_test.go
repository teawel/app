package options

import "testing"

func TestBoolOption_AsHTML(t *testing.T) {
	option := NewCheckbox("Is Checked", "isChecked")
	option.Value = 1
	t.Log(option)
}
