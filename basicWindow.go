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
	vertical bool
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

// Returns true if initialized
func (w *BasicWindow) Initialized() bool {
	if w == nil {return false}
	if w.parent == nil {return false}
	return true
}

// Returns true if the status is valid
func (w *BasicWindow) Ready() bool {
	if ! w.Initialized() {return false}
	if ! w.parent.Ready() {return false}
	return w.ready
}

func (w *BasicWindow) Box() *gui.Node {
	if ! w.Initialized() {return nil}
	if (w.win == nil) {
		w.Draw()
	}
	return w.box
}

func (w *BasicWindow) Vertical() {
	log.Log(INFO, "BasicWindow() w.vertical =", w.vertical)
	if ! w.Initialized() {
		log.Warn("BasicWindow() not Initialized yet()")
		return
	}
	if w.Ready() {
		log.Warn("BasicWindow() is already created. You can not change it to Vertical now (TODO: fix this)")
		return
	}
	w.vertical = true
	log.Log(INFO, "BasicWindow() w.vertical =", w.vertical)
}

func (w *BasicWindow) Make() {
	if ! w.Initialized() {return}
	// various timeout settings
	w.win = w.parent.RawWindow(w.name)
	w.win.Custom = func() {
		log.Warn("BasicWindow.Custom() closed. TODO: handle this", w.name)
	}
	if w.vertical {
		w.box = w.win.NewBox("bw vbox", false)
		log.Log(INFO, "BasicWindow.Custom() made vbox")
	} else {
		w.box = w.win.NewBox("bw hbox", true)
		log.Log(INFO, "BasicWindow.Custom() made hbox")
	}

	w.ready = true
}

func (w *BasicWindow) Draw() {
	if ! w.Initialized() {return}
	// various timeout settings
	w.win = w.parent.NewWindow(w.name)
	w.win.Custom = func() {
		log.Warn("BasicWindow.Custom() closed. TODO: handle this", w.name)
	}
	if w.vertical {
		w.box = w.win.NewBox("bw vbox", false)
		log.Log(INFO, "BasicWindow.Custom() made vbox")
	} else {
		w.box = w.win.NewBox("bw hbox", true)
		log.Log(INFO, "BasicWindow.Custom() made hbox")
	}

	w.ready = true
}


func NewBasicWindow(parent *gui.Node, name string) *BasicWindow {
	var w *BasicWindow
	w = &BasicWindow {
		parent: parent,
		name: name,
		vertical: false,
	}

	return w
}
