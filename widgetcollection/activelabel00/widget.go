package activelabel00

import (
	"syscall/js"

	pexu_dom "github.com/AnimusPEXUS/wasmtools/dom"
	"github.com/AnimusPEXUS/wasmtools/elementtreeconstructor"
)

type ActiveLabel00 struct {
	Element *elementtreeconstructor.ElementMutator
}

func NewActiveLabel(
	text string,
	title string,

	onclick func(),

	etc *elementtreeconstructor.ElementTreeConstructor,

) *ActiveLabel00 {

	self := &ActiveLabel00()

	el := etc.CreateElement("a").
		ExternalUse(applyAStyle).
		AssignSelf(&self.add_new_proxy_target_button).
		AppendChildren(
			etc.CreateTextNode(text),
		)

	if title != "" {
		el.Set("title", "add new proxy target")
	}

	return self
}

func applyAStyle(ed *elementtreeconstructor.ElementMutator) {
	ed.
		SetStyle("color", "blue").
		SetStyle("cursor", "pointer").
		SetStyle("text-decoration", "underline")
}
