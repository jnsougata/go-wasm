package dom

import "syscall/js"

type buttonTag struct {
	JSValue Value
}

func (b *buttonTag) Attributes(attributes ...Attribute) {
	for _, attr := range attributes {
		key, value := attr.merge()
		b.JSValue.Set(key, value)
	}
}

func (b *buttonTag) OnClick(f func(this Value, args []Value) interface{}) {
	b.JSValue.Set("onclick", js.FuncOf(f))
}
