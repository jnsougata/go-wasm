package dom

type Image struct {
	Height int
	Width  int
	Src    string
}

type imgTag struct {
	JSValue Value
}

func (i *imgTag) SetAttributes(attributes ...Attribute) {
	for _, attr := range attributes {
		key, value := attr.merge()
		i.JSValue.Set(key, value)
	}
}
