package dom

import (
	"syscall/js"
)

type Document struct {
	JSValue *js.Value
}

func NewDocumentFromJsValue(jsvalue *js.Value) *Document {
	self := &Document{jsvalue}
	return self
}

func (self *Document) CreateElementNS(ns string, name string) *Element {
	return &Element{&Node{&[]js.Value{self.JSValue.Call("createElementNS", ns, name, js.Undefined())}[0]}}
}

func (self *Document) CreateElement(name string) *Element {
	return &Element{&Node{&[]js.Value{self.JSValue.Call("createElement", name, js.Undefined())}[0]}}

}

func (self *Document) NewTextNode(text string) *Node {
	return &Node{&[]js.Value{self.JSValue.Call("createTextNode", text)}[0]}
}

func (self *Document) GetBody() *Element {
	return &Element{&Node{&[]js.Value{self.JSValue.Get("body")}[0]}}
}
