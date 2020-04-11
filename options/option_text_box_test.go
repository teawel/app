package options

import "testing"

func TestMemoOption_AsHTML(t *testing.T) {
	option := NewTextBox("Description", "description")
	t.Log(option)
}
