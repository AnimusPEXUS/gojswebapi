package dom

import (
	"syscall/js"
)

type Document struct {
	js.Value
}

func NewDocumentFromJsValue(jsvalue js.Value) *Document {
	self := &Document{jsvalue}
	return self
}

func (self *Document) CreateElementNS(ns string, name string) *Element {
	return &Element{Node{self.Value.Call("createElementNS", ns, name, js.Undefined())}}
}

func (self *Document) CreateElement(name string) *Element {
	return &Element{Node{self.Value.Call("createElement", name, js.Undefined())}}

}

func (self *Document) NewTextNode(text string) *Node {
	return &Node{self.Value.Call("createTextNode", text)}
}

func (self *Document) GetBody() *Element {
	return &Element{Node{self.Get("body")}}
}
