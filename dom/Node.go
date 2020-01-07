package dom

import (
	"syscall/js"
)

type Node struct {
	js.Value
}

func (self *Node) AppendChild(node *Node) *Node {
	return &Node{self.Call("appendChild", node)}
}

func (self *Node) GetFirstChild() *Node {

	ret := (*Node)(nil)

	r := self.Value.Get("firstChild")
	if r != js.Null() {
		ret = &Node{r}
	}

	return ret
}

func (self *Node) RemoveChild(c *Node) *Node {
	return &Node{self.Value.Call("removeChild", c.Value)}
}

// func (self *Node) RemoveAllChildren() {
// 	self.Value.Call("removeAllChildren")
// }
