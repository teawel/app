package options

import "testing"

func TestMemoOption_AsHTML(t *testing.T) {
	option := NewMemoOption("Description", "description")
	t.Log(option.AsHTML())
}
