package js

import "syscall/js"

type buttonTag struct {
	JSValue js.Value
}

func (b *buttonTag) SetAttribute(key, value string) {
	b.JSValue.Set(key, value)
}

func (b *buttonTag) OnCLick(f func(this Value, args []Value) interface{}) {
	b.JSValue.Set("onclick", js.FuncOf(f))
}
