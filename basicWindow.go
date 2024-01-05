/*
	A Standard Window
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

type BasicWindow struct {
	ready	bool
	hidden	bool
	name	string

	parent	*gui.Node
	win	*gui.Node	// window widget
	box	*gui.Node	// box

	Custom func()
}

func (w *BasicWindow) Hide() {
	if ! w.Ready() {return}
	w.win.Hide()
	w.hidden = true
	return
}

func (w *BasicWindow) Show() {
	if ! w.Ready() {return}
	w.win.Show()
	w.hidden = false
	return
}

func (w *BasicWindow) Toggle() {
	if ! w.Ready() {return}
	if w.hidden {
		w.Show()
		w.hidden = false
	} else {
		w.Hide()
		w.hidden = true
	}
	return
}

func (w *BasicWindow) Title(title string) {
	if ! w.Ready() {return}
	w.win.SetText(title)
	return
}

// Returns true if the status is valid
func (w *BasicWindow) Ready() bool {
	if w == nil {return false}
	if w.parent == nil {return false}
	if ! w.parent.Ready() {return false}
	if (w.win == nil) {
		w.Draw()
	}
	return w.ready
}

func (w *BasicWindow) Box() *gui.Node {
	return w.box
}

func (w *BasicWindow) Draw() {
	w.ready = true
	return
}


func NewBasicWindow(parent *gui.Node, name string) *BasicWindow {
	var w *BasicWindow
	w = &BasicWindow {
		parent: parent,
		name: name,
	}

	// various timeout settings
	w.win = w.parent.NewWindow(name)
	w.win.Custom = func() {
		log.Println("BasicWindow.Custom() closed. TODO: handle this", w.name)
	}
	w.box = w.win.NewBox("hBox", true)

	return w
}
