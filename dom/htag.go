package dom

type hTag struct {
	JSValue Value
}

func (h *hTag) Attributes(attributes ...Attribute) {
	for _, attr := range attributes {
		key, value := attr.merge()
		h.JSValue.Set(key, value)
	}
}
