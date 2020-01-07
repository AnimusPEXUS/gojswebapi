package elementtreeconstructor

import (
	"github.com/AnimusPEXUS/wasmtools/dom"
)

type ElementTreeConstructor struct {
	document *dom.Document
}

func NewElementTreeConstructor(document *dom.Document) *ElementTreeConstructor {
	self := &ElementTreeConstructor{
		document: document,
	}
	return self
}

func (self *ElementTreeConstructor) CreateTextNode(
	text string,
) *dom.Node {
	return self.document.NewTextNode(text)
}

func (self *ElementTreeConstructor) CreateElement(
	name string,
	attributes []*dom.Attribute,
	children []*dom.Node,
) *dom.Element {
	return self.CreateElementNS(nil, name, attributes, children)
}

func (self *ElementTreeConstructor) CreateElementNS(
	namespace *string,
	name string,
	attributes []*dom.Attribute,
	children []*dom.Node,
) *dom.Element {

	var ret *dom.Element

	if namespace != nil {
		ret = self.document.CreateElementNS(*namespace, name)
	} else {
		ret = self.document.CreateElement(name)
	}

	for _, i := range attributes {
		if i.Namespace == nil {
			ret.SetAttribute(i.Name, i.Value)
		} else {
			ret.SetAttributeNS(*i.Namespace, i.Name, i.Value)
		}
	}

	for _, i := range children {
		ret.Append(i)
	}

	return ret
}

func (self *ElementTreeConstructor) ReplaceChildren(new_children []*dom.Node) {

	n := &dom.Node{self.document.Value}

	for i := n.GetFirstChild(); i != nil; i = n.GetFirstChild() {
		n.RemoveChild(i)
	}

	for _, i := range new_children {
		// log.Println("appending child", i, i.Value)
		n.AppendChild(i)
	}

}
