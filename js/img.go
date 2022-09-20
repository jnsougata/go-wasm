package js

import "syscall/js"

type Image struct {
	Height int
	Width  int
	Src    string
}

type imgTag struct {
	JSValue js.Value
}

func (i *imgTag) SetAttribute(key, value string) {
	i.JSValue.Set(key, value)
}
