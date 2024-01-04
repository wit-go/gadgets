package logsettings

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

func (d *LogSettings) Show() {
	if ! d.Ready() {
		log.Warn("LogSettings.Show() window is not Ready()")
		return
	}
	log.Warn("LogSettings.Show() window")
	if d.hidden {
		log.Warn("LogSettings.Show() window HERE window =", d.window)
		if d.window == nil {
			log.Warn("LogSettings.Show() create the window")
			d.draw()
		}
		d.window.Show()
	}
	d.hidden = false
}

func (d *LogSettings) Hide() {
	if ! d.Ready() {
		log.Warn("LogSettings.Show() window is not Ready()")
		return
	}
	log.Warn("LogSettings.Hide() window")
	if ! d.hidden {
		d.window.Hide()
	}
	d.hidden = true
}

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func (d *LogSettings) draw() {
	if ! d.Ready() {return}
	var newW, newB, g *gui.Node

	newW = d.parent.NewWindow("Debug Flags")
	newW.Custom = d.parent.StandardClose

	newB = newW.NewBox("hBox", true)
	g = newB.NewGroup("Show").Pad()

	g.NewButton("log.SetTmp()", func () {
		log.SetTmp()
	})

	g.NewButton("log.SetAll(true)", func () {
		log.SetAll(true)
	})

	g.NewButton("log.SetAll(false)", func () {
		log.SetAll(false)
	})

	g.NewButton("Dump Flags", func () {
		// ShowDebugValues()
		log.ShowFlags()
	})

	g = newB.NewGroup("List")
	g = g.NewGrid("flags grid", 5, 2)
	flags := log.ShowFlags()
	for _, f := range flags {
		log.Log(true, "Get() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
		gadgets.NewLogFlag(g, f)
	}
}
