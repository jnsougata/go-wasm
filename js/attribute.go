package js

import "fmt"

type Attribute struct {
	Name   string
	Value  string
	Values map[string]any
}

func (attr *Attribute) merge() (string, string) {
	if attr.Value != "" {
		return attr.Name, attr.Value
	}
	var value string
	for k, v := range attr.Values {
		value += k + ":" + fmt.Sprintf("%v", v) + ";"
	}
	return attr.Name, value
}

type Style struct {
	Object string
	Value  string
	Values []string
}

func (style *Style) merge() (string, string) {
	if style.Value != "" {
		style.Values = append(style.Values, style.Value)
	}
	if len(style.Values) == 0 {
		panic("Style values cannot be empty")
	}
	if style.Object == "" {
		panic("Style object cannot be empty")
	}
	var value string
	for _, v := range style.Values {
		value += v
	}
	return style.Object, value
}
