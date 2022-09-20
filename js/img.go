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

func (i *imgTag) SetAttribute(attribute Attribute) {
	key, value := attribute.merge()
	i.JSValue.Set(key, value)
}
