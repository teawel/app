package options

import "testing"

func TestStringOption_AsHTML(t *testing.T) {
	option := NewStringOption("Name", "name")
	option.MaxLength = 100
	option.Attr("style", "color: red")
	t.Log(option.AsHTML())
}
