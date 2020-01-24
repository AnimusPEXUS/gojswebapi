package elementtreeconstructor

import (
	"syscall/js"

	"github.com/AnimusPEXUS/wasmtools/dom"
)

type ElementMutator struct {
	Element *dom.Element
}

func NewElementMutatorNS(doc *dom.Document, ns string, name string) *ElementMutator {
	ret := &ElementMutator{doc.CreateElementNS(ns, name)}
	return ret
}

func NewElementMutator(doc *dom.Document, name string) *ElementMutator {
	ret := &ElementMutator{doc.CreateElement(name)}
	return ret
}

func NewElementMutatorFromElement(e *dom.Element) *ElementMutator {
	ret := &ElementMutator{e}
	return ret
}

func (self *ElementMutator) AppendChildren(children ...interface{}) *ElementMutator {

	// FIXME: append and remove operations have to be done at Node level

	for _, i := range children {

		switch i.(type) {
		default:
			panic("unsupported type")
		case *ElementMutator:
			self.Element.Append(i.(*ElementMutator).Element.AsNode())
		case *dom.Element:
			self.Element.Append(i.(*dom.Element).AsNode())
		case *dom.Node:
			self.Element.Append(i.(*dom.Node))
		}

	}
	return self
}

func (self *ElementMutator) RemoveChildren() *ElementMutator {
	self.Element.Node.RemoveAllChildren()
	return self
}

func (self *ElementMutator) Remove(children ...interface{}) *ElementMutator {

	// FIXME: append and remove operations have to be done at Node level

	for _, i := range children {
		var n *dom.Node
		switch i.(type) {
		default:
			panic("unsupported type")
		case *ElementMutator:
			n = i.(*ElementMutator).Element.AsNode()
		case *dom.Element:
			n = i.(*dom.Element).AsNode()
		case *dom.Node:
			n = (i.(*dom.Node))
		}

		self.Element.RemoveChild(n)

	}
	return self
}

func (self *ElementMutator) RemoveFromParent() *ElementMutator {
	self.Parent().Remove(self)
	return self
}

func (self *ElementMutator) Parent() *ElementMutator {

	// FIXME: this is Node level operation

	ret := (*ElementMutator)(nil)
	t := self.Element.ParentElement()
	if t != nil {
		ret = NewElementMutatorFromElement(t)
	}
	return ret
}

func (self *ElementMutator) AssignSelf(variable **ElementMutator) *ElementMutator {
	*variable = self
	return self
}

func (self *ElementMutator) AssignSelfDom(variable **dom.Element) *ElementMutator {
	*variable = self.Element
	return self
}

func (self *ElementMutator) ExternalUse(cb func(*ElementMutator)) *ElementMutator {
	cb(self)
	return self
}

func (self *ElementMutator) Call(property string, ret *interface{}, args ...interface{}) *ElementMutator {
	t := self.Element.Node.Value.Call(property, args...)
	if ret != nil {
		*ret = t
	}
	return self
}

func (self *ElementMutator) SetAttribute(name string, value string) *ElementMutator {
	self.Element.SetAttribute(name, value)
	return self
}

// func (self *ElementMutator) GetAttribute((name string, value string, ret *interface{})) *ElementMutator {
// 	self.Element.SetAttribute(name, value)
// 	if ret != nil {
// 		*ret = t
// 	}
// 	return self
// }
//

func (self *ElementMutator) Set(property string, value interface{}) *ElementMutator {
	self.Element.Node.Value.Set(property, value)
	return self
}

func (self *ElementMutator) Get(property string) interface{} {
	return self.Element.Node.Value.Get(property)
}

func (self *ElementMutator) GetJsValue(property string) js.Value {
	return self.Element.Node.Value.Get(property)
}

func (self *ElementMutator) GetAssign(property string, ret *interface{}) *ElementMutator {
	t := self.Element.Node.Value.Get(property)
	if ret != nil {
		*ret = t
	}
	return self
}

func (self *ElementMutator) SetStyle(property string, value interface{}) *ElementMutator {
	self.Element.Node.Value.Get("style").Set(property, value)
	return self
}

func (self *ElementMutator) AddListener(
	event string,
	cb func(this js.Value, args []js.Value) interface{},
) *ElementMutator {
	self.Element.Node.Value.Get(event).Call("addListener", js.FuncOf(cb))
	return self
}

func (self *ElementMutator) AddEventListener(
	event string,
	cb func(this js.Value, args []js.Value) interface{},
) *ElementMutator {
	self.Element.Node.Value.Call("addEventListener", event, js.FuncOf(cb))
	return self
}

func (self *ElementMutator) ToDomElement() *dom.Element {
	return self.Element
}

func (self *ElementMutator) ToDomNode() *dom.Node {
	return self.ToDomElement().AsNode()
}
