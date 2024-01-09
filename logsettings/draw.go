package logsettings

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

func (d *LogSettings) Show() {
	if ! d.Ready() { return }
	d.win.Show()
}

func (d *LogSettings) Hide() {
	if ! d.Ready() { return }
	d.win.Hide()
}

// alternates between showing and hiding the window
func (d *LogSettings) Toggle() {
	if ! d.Ready() { return }
	d.win.Toggle()
}

// Let's you toggle on and off the various types of debugging output
// These checkboxes should be in the same order as the are printed
func (d *LogSettings) draw() {
	var g *gui.Node

	d.win = gadgets.NewBasicWindow(d.parent, "Debug Flags")
	g = d.win.Box().NewGroup("Show").Pad()
	g = g.NewBox("bw vbox", false)
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
			log.Info("Dump Flags", s)
			for _, f := range wg.flags {
				log.Info("Dump Flags\t", f.Get(), f.Name, ":", f.Desc)
			}
		}
	})

	d.flagG = d.win.Box().NewGroup("Subsystem (aka package)")
	d.flagG = d.flagG.NewBox("bw vbox", false)

	g.NewButton("Add all Flags", func () {
		flags := log.ShowFlags()
		for _, f := range flags {
			addFlag(d.flagG, f)
		}
	})

	g.NewButton("Close", func () {
		d.Hide()
	})

	flags := log.ShowFlags()
	for _, f := range flags {
		addFlag(d.flagG, f)
	}
}

func addFlag(p *gui.Node, newf *log.LogFlag) {
	var flagWidgets *flagGroup
	if newf == nil { return }
	if p == nil { return }

	subsys := newf.GetSubsystem()
	name := newf.GetName()

	if myLogGui.groups[subsys] == nil {
		flagWidgets = new(flagGroup)
		flagWidgets.parent = p
		flagWidgets.name = subsys
		flagWidgets.group = p.NewGroup(subsys)
		flagWidgets.grid = flagWidgets.group.NewGrid("flags grid", 3, 1)
		myLogGui.groups[subsys] = flagWidgets
	} else {
		flagWidgets = myLogGui.groups[subsys]
	}

	for _, f := range flagWidgets.flags {
		if f.Name == name {
			log.Info("addFlag() FOUND FLAG", f)
			return
		}
	}
	log.Info("addFlag() Adding new flag:", subsys, name)
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
