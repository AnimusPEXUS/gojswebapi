package select00

import (
	"syscall/js"

	pexu_dom "github.com/AnimusPEXUS/wasmtools/dom"
	"github.com/AnimusPEXUS/wasmtools/dom/elementtreeconstructor"
)

type Select00 struct {
	document *pexu_dom.Document

	value_select *elementtreeconstructor.ElementMutator

	Value   string
	Element *pexu_dom.Element
}

func NewSelect00(
	document *pexu_dom.Document,
	values [][2]string,
	preselected string,
	onchange func(),
) *Select00 {

	self := &Select00{}

	self.document = document

	etc := elementtreeconstructor.NewElementTreeConstructor(document)

	self.value_select = etc.CreateElement("select")

	for _, i := range values {
		self.AppendOption(i[0], i[1])
	}

	if preselected != "" {
		self.value_select.Set("value", preselected)
	}

	self.value_select.Set(
		"onchange",
		js.FuncOf(
			func(
				this js.Value,
				args []js.Value,
			) interface{} {
				self.Value = self.value_select.GetJsValue("value").String()
				onchange()
				return false
			},
		),
	)

	self.Element = self.value_select.Element

	return self
}

func (self *Select00) AppendOption(key, value string) {
	etc := elementtreeconstructor.NewElementTreeConstructor(self.document)
	self.value_select.
		AppendChildren(
			etc.
				CreateElement("option").
				Set("value", key).
				AppendChildren(
					etc.CreateTextNode(value),
				),
		)
}

func (self *Select00) SetSelected(value string) {
	self.value_select.Set("value", value)
}
