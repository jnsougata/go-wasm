package dom

type divTag struct {
	JSValue Value
}

func (d *divTag) Attributes(attributes ...Attribute) {
	for _, attr := range attributes {
		key, value := attr.merge()
		d.JSValue.Set(key, value)
	}
}

func (d *divTag) Style() {

}

func (d *divTag) Append(children ...Value) {
	for _, child := range children {
		d.JSValue.Call("appendChild", child)
	}
}
