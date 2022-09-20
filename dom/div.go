package dom

import "syscall/js"

type divTag struct {
	JSValue js.Value
}

func (d *divTag) Attributes(attributes ...Attribute) {
	for _, attr := range attributes {
		key, value := attr.merge()
		d.JSValue.Set(key, value)
	}
}

func (d *divTag) Append(children ...Value) {
	for _, child := range children {
		d.JSValue.Call("appendChild", child)
	}
}
