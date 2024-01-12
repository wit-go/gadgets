/*
	A Labeled label:

	-----------------------------
	|            |              |
	|   Food:    |    Apple     |
	|            |              |
	-----------------------------
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

type OneLiner struct {
	p	*gui.Node	// parent widget
	l	*gui.Node	// label widget
	v	*gui.Node	// value widget

	value	string
	label	string

	Custom func()
}

func (n *OneLiner) Get() string {
	return n.value
}

func (n *OneLiner) Set(value string) *OneLiner {
	log.Log(INFO, "OneLiner.Set() =", value)
	if (n.v != nil) {
		n.v.Set(value)
	}
	n.value = value
	return n
}

func (n *OneLiner) SetLabel(value string) *OneLiner {
	log.Log(INFO, "OneLiner.SetLabel() =", value)
	if (n.l != nil) {
		n.l.Set(value)
	}
	return n
}

func (n *OneLiner) Enable() {
	log.Log(INFO, "OneLiner.Enable()")
	if (n.v != nil) {
		n.v.Show()
	}
}

func (n *OneLiner) Disable() {
	log.Log(INFO, "OneLiner.Disable()")
	if (n.v != nil) {
		n.v.Hide()
	}
}

func NewOneLiner(n *gui.Node, name string) *OneLiner {
	d := OneLiner {
		p: n,
		value: "",
	}

	// various timeout settings
	d.l = n.NewLabel(name)
	d.v = n.NewLabel("")
	d.v.Custom = func() {
		d.value = d.v.GetText()
		log.Log(INFO, "OneLiner.Custom() user changed value to =", d.value)
	}

	return &d
}
