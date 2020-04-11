package options

import "testing"

func TestStringOption_AsHTML(t *testing.T) {
	option := NewTextField("Name", "name")
	option.MaxLength = 100
	option.Attr("style", "color: red")
	t.Log(option)
}
