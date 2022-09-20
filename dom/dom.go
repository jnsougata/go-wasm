package dom

import (
	"fmt"
	"syscall/js"
)

type Value = js.Value

func New() *document {
	d := &document{}
	d.fetch()
	return d
}

type document struct {
	inner js.Value
	Body  js.Value
	Head  js.Value
}

func (d *document) fetch() {
	d.inner = js.Global().Get("document")
	d.Head = d.inner.Get("head")
	d.Body = d.inner.Get("body")
}

func (d *document) Append(children ...Value) {
	for _, c := range children {
		d.Body.Call("appendChild", c)
	}
}

func (d *document) StyleSheet(path string) {
	link := d.inner.Call("createElement", "link")
	link.Set("rel", "stylesheet")
	link.Set("href", path)
	link.Set("type", "text/css")
	d.Head.Call("appendChild", link)
}

func (d *document) Styles(style ...Style) {
	var stl string
	for _, v := range style {
		key, value := v.merge()
		stl += fmt.Sprintf("%s {%s}", key, value)
	}
	s := d.inner.Call("createElement", "style")
	s.Set("innerHTML", stl)
	d.Head.Call("appendChild", s)
}

func (d *document) NewImg(image *Image) *imgTag {
	i := d.inner.Call("createElement", "img")
	if image != nil {
		if image.Height != 0 {
			i.Set("height", image.Height)
		}
		if image.Width != 0 {
			i.Set("width", image.Width)
		}
		if image.Src != "" {
			i.Set("src", image.Src)
		}
		return &imgTag{JSValue: i}
	}
	return nil
}

func (d *document) NewHTag(text string, tag int) *hTag {
	h3 := d.inner.Call("createElement", fmt.Sprintf("h%d", tag))
	h3.Set("innerHTML", text)
	return &hTag{JSValue: h3}
}

func (d *document) NewButton(text string) *buttonTag {
	b := d.inner.Call("createElement", "button")
	b.Set("innerHTML", text)
	return &buttonTag{JSValue: b}
}

func (d *document) NewDiv() *divTag {
	div := d.inner.Call("createElement", "div")
	return &divTag{JSValue: div}
}
