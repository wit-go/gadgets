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
		log.Warn("LogSettings.Hide() window is not Ready()")
		return
	}
	log.Warn("LogSettings.Hide() window")
	d.window.Hide()
	d.hidden = true
}

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func (d *LogSettings) draw() {
	if ! d.Ready() {return}
	var g *gui.Node

	d.window = d.parent.NewWindow("Debug Flags")
	d.window.Custom = d.parent.StandardClose

	d.box = d.window.NewBox("hBox", true)
	g = d.box.NewGroup("Show").Pad()
	d.buttonG = g

	g.NewButton("Redirect STDOUT to /tmp/", func () {
		log.SetTmp()
	})

	g.NewButton("restore defaults", func () {
		for _, wg := range myLogGui.groups {
			for _, f := range wg.flags {
				f.SetDefault()
			}
		}
	})

	g.NewButton("all on", func () {
		for _, wg := range myLogGui.groups {
			for _, f := range wg.flags {
				f.Set(true)
			}
		}
	})

	g.NewButton("all off", func () {
		for _, wg := range myLogGui.groups {
			for _, f := range wg.flags {
				f.Set(false)
			}
		}
	})

	g.NewButton("Dump Flags", func () {
		// ShowDebugValues()
		log.ShowFlags()
		for s, wg := range myLogGui.groups {
			log.Log(true, "Dump Flags", s)
			for _, f := range wg.flags {
				log.Log(true, "Dump Flags\t", f.Name, ":", f.Desc)
			}
		}
	})

	d.flagG = d.box.NewGroup("Subsystem (aka package)")

	g.NewButton("Add all Flags", func () {
		flags := log.ShowFlags()
		for _, f := range flags {
			log.Log(true, "Get() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
			addFlag(d.flagG, f)
		}
	})

	g.NewButton("Close", func () {
		d.Hide()
	})

	flags := log.ShowFlags()
	for _, f := range flags {
		log.Log(true, "Get() ", "(" + f.Subsystem + ")", f.Name, "=", f.B, ":", f.Desc)
		addFlag(d.flagG, f)
	}
}

func addFlag(p *gui.Node, newf *log.LogFlag) {
	var flagWidgets *flagGroup
	if newf == nil { return }
	if p == nil { return }

	if myLogGui.groups[newf.Subsystem] == nil {
		flagWidgets = new(flagGroup)
		flagWidgets.parent = p
		flagWidgets.name = newf.Subsystem
		flagWidgets.group = p.NewGroup(newf.Subsystem)
		flagWidgets.grid = flagWidgets.group.NewGrid("flags grid", 3, 1)
		myLogGui.groups[newf.Subsystem] = flagWidgets
	} else {
		flagWidgets = myLogGui.groups[newf.Subsystem]
	}

	for _, f := range flagWidgets.flags {
		log.Warn("addFlag() Already has flag =", f)
		if f.Name == newf.Name {
			log.Warn("addFlag() FOUND FLAG", f)
			return
		}
	}
	newWidget := gadgets.NewLogFlag(flagWidgets.grid, newf)
	flagWidgets.flags = append(flagWidgets.flags, newWidget)
}

type flagGroup struct {
	name	string // should be set to the flag.Subsystem

	parent	*gui.Node // where to draw our group
	group *gui.Node
	grid *gui.Node

	// the widget for each flag
	flags []*gadgets.LogFlag
}
