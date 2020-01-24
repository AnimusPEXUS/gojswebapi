package dom

// import (
// 	"syscall/js"
// )

type Element struct {
	// js.Value
	Node
}

func (self *Element) Append(node *Node) {
	self.Call("append", node)
}

func (self *Element) Remove() {
	// NOTE: this removes Element it self from it's parent
	self.Call("remove")
}

func (self *Element) RemoveChild(node *Node) {
	// NOTE: this removes Element it self from it's parent
	self.Node.RemoveChild(node)
}

func (self *Element) ParentNode() *Node {
	return self.Node.ParentNode()
}

func (self *Element) ParentElement() *Element {
	return self.Node.ParentElement()
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

// type Attribute struct {
// 	Namespace *string
// 	Name      string
// 	Value     string
// }
