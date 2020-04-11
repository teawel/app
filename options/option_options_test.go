package options

import "testing"

func TestOptions_AsHTML(t *testing.T) {
	options := NewOptions("Select Book", "book")
	options.Attr("style", "color: red")
	options.AddValue("PHP", "php")
	options.AddValue("Java", "java")
	options.AddValue("Golang", "golang")
	t.Log(options)
}
