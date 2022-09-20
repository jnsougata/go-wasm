package js

import "syscall/js"

type hTag struct {
	JSValue js.Value
}

func (h *hTag) SetAttribute(attribute Attribute) {
	key, value := attribute.merge()
	h.JSValue.Set(key, value)
}
