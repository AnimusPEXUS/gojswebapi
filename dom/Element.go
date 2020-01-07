package dom

import (
	"syscall/js"
)

type Element struct {
	js.Value
}

func (self *Element) Append(node *Node) {
	self.Call("append", node)
}

func (self *Element) AsNode() *Node {
	return &Node{self.Value}
}

func (self *Element) SetAttribute(name string, value string) {
	self.Call("setAttribute", name, value)
}

func (self *Element) SetAttributeNS(namespace string, name string, value string) {
	self.Call("setAttributeNS", namespace, name, value)
}

type Attribute struct {
	Namespace *string
	Name      string
	Value     string
}
