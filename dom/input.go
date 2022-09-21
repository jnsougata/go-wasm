package dom

type inputTag struct {
	JSValue Value
}

func (i *inputTag) Attributes(attributes ...Attribute) {
	for _, attr := range attributes {
		key, value := attr.merge()
		i.JSValue.Set(key, value)
	}
}
