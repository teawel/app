package options

import "testing"

func TestBoolOption_AsHTML(t *testing.T) {
	option := NewBoolOption("Is Checked", "isChecked")
	option.Value = 1
	t.Log(option.AsHTML())
}
