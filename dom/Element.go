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
