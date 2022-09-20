package js

import "syscall/js"

type hTag struct {
	JSValue js.Value
}

func (h *hTag) SetAttribute(key, value string) {
	h.JSValue.Set(key, value)
}
