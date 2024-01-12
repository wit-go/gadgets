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
	title	string

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
	log.Warn("BasicWindow() Vertical() START w.vertical =", w.vertical, w.title)
	if w == nil {return}
	w.vertical = true
	log.Warn("BasicWindow() Vertical() END w.vertical =", w.vertical, w.title)
}

func (w *BasicWindow) Horizontal() {
	log.Warn("BasicWindow() Horizontal() START w.vertical =", w.vertical, w.title)
	if w == nil {return}
	w.vertical = false
	log.Warn("BasicWindow() Horizontal() END w.vertical =", w.vertical, w.title)
}

func (w *BasicWindow) Make() {
	if ! w.Initialized() {return}
	// various timeout settings
	w.win = w.parent.RawWindow(w.title)
	w.win.Custom = func() {
		log.Warn("BasicWindow.Custom() closed. TODO: handle this", w.title)
	}
	if w.vertical {
		w.box = w.win.NewVerticalBox("BW VBOX")
		log.Log(INFO, "BasicWindow.Make() made NewVerticalBox", w.title)
	} else {
		w.box = w.win.NewHorizontalBox("BW HBOX")
		log.Log(INFO, "BasicWindow.Make() made NewHorizontalBox", w.title)
	}

	w.ready = true
}

func (w *BasicWindow) Draw() {
	if ! w.Initialized() {return}
	// various timeout settings
	w.win = w.parent.NewWindow(w.title)
	w.win.Custom = func() {
		log.Warn("BasicWindow.Custom() closed. TODO: handle this", w.title)
	}
	log.Warn("BasicWindow.Draw() about to make a box vertical =", w.vertical, w.title)
	if w.vertical {
		w.box = w.win.NewVerticalBox("BW VBOX")
		log.Log(INFO, "BasicWindow.Draw() made vbox title =", w.title)
	} else {
		w.box = w.win.NewHorizontalBox("BW HBOX")
		log.Log(INFO, "BasicWindow.Draw() made hbox title =", w.title)
	}

	w.ready = true
}


func NewBasicWindow(parent *gui.Node, title string) *BasicWindow {
	var w *BasicWindow
	w = &BasicWindow {
		parent: parent,
		title: title,
		vertical: false,
	}
	log.Warn("NewBasicWindow() END")

	return w
}
