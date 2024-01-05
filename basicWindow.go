/*
	A Standard Window
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

type BasicWindow struct {
	hidden	bool
	name	string

	p	*gui.Node	// parent widget
	win	*gui.Node	// window widget
	box	*gui.Node	// box

	Custom func()
}

func (w *BasicWindow) Hide() {
	w.win.Hide()
	w.hidden = true
	return
}

func (w *BasicWindow) Show() {
	w.win.Show()
	w.hidden = false
	return
}

func (w *BasicWindow) Toggle() {
	if w.hidden {
		w.Show()
		w.hidden = false
	} else {
		w.Hide()
		w.hidden = true
	}
	return
}

func (w *BasicWindow) Box() *gui.Node {
	return w.box
}

func NewBasicWindow(parent *gui.Node, name string) *BasicWindow {
	var w *BasicWindow
	w = &BasicWindow {
		p: parent,
		name: name,
	}

	// various timeout settings
	w.win = w.p.NewWindow(name)
	w.win.Custom = func() {
		log.Println("BasicWindow.Custom() closed. TODO: handle this", w.name)
	}
	w.box = w.win.NewBox("hBox", true)

	return w
}
