package dom

import (
	"syscall/js"
)

type ToNodeConvertable interface {
	AsNode() *Node
}

type Node struct {
	js.Value
}

func (self *Node) AppendChild(node *Node) *Node {
	return &Node{self.Call("appendChild", node)}
}

func (self *Node) GetFirstChild() *Node {

	ret := (*Node)(nil)

	r := self.Value.Get("firstChild")
	if !r.IsNull() {
		ret = &Node{r}
	}

	return ret
}

func (self *Node) RemoveChild(c *Node) *Node {
	return &Node{self.Value.Call("removeChild", c.Value)}
}

func (self *Node) ParentNode() *Node {
	t := self.Get("parentNode")
	if t.IsNull() || t.IsUndefined() {
		return nil
	}
	return &Node{t}
}

func (self *Node) ParentElement() *Element {
	t := self.Get("parentElement")
	if t.IsNull() || t.IsUndefined() {
		return nil
	}
	return &Element{Node{t}}
}

func (self *Node) AsNode() *Node {
	return self
}

func (self *Node) RemoveAllChildren() {
	for a := self.GetFirstChild(); a != nil; a = self.GetFirstChild() {
		self.RemoveChild(a)
	}
}
