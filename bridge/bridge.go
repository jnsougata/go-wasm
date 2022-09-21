package bridge

import (
	"fmt"
	"syscall/js"
)

type Arg = js.Value

var Invoker = js.Global()

type Class struct {
	Name       string
	Attributes map[string]string
}

type DomObject struct {
	root Arg
}

func (d *DomObject) OnClick(listener func(this Arg, args []Arg) interface{}) {
	d.root.Set("onclick", js.FuncOf(listener))
}

func (d *DomObject) Set(attribute string, value string) {
	d.root.Set(attribute, value)
}

func (d *DomObject) Get(attribute string) string {
	return d.root.Get(attribute).String()
}

func (d *DomObject) Call(JSMethod string, object *DomObject) {
	d.root.Call(JSMethod, object.root)
}

type Head struct {
	root Arg
}

func (h *Head) AppendChild(dom *DomObject) {
	h.root.Call("appendChild", dom.root)
}

type Body struct {
	root Arg
}

func (b *Body) AppendChild(dom *DomObject) {
	b.root.Call("appendChild", dom.root)
}

func New() *Bridge {
	root := Invoker.Get("document")
	head := &Head{root.Get("head")}
	body := &Body{root.Get("body")}
	return &Bridge{
		root: root,
		Head: head,
		Body: body,
	}
}

type Bridge struct {
	root Arg
	Body *Body
	Head *Head
}

func (d *Bridge) StyleSheet(path string) {
	link := d.CreateElement("link")
	link.Set("rel", "stylesheet")
	link.Set("href", path)
	link.Set("type", "text/css")
	d.Head.root.Call("appendChild", link.root)
}

func (d *Bridge) CSSClasses(classes ...Class) {
	var styleString string
	for _, class := range classes {
		if class.Name != "" {
			var attrs string
			for k, v := range class.Attributes {
				attrs += fmt.Sprintf("%s: %s; ", k, v)
			}
			styleString += fmt.Sprintf(" %s {%s} ", class.Name, attrs)
		}
	}
	tag := d.CreateElement("style")
	tag.Set("innerHTML", styleString)
	d.Head.AppendChild(tag)
}

func (d *Bridge) CreateElement(name string) *DomObject {
	i := d.root.Call("createElement", name)
	return &DomObject{i}
}
