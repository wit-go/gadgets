package logsettings

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

// This initializes the main object
// You can only have one of these
func New(p *gui.Node) *LogSettings {
	if myLogGui != nil {return myLogGui}
	myLogGui = new(LogSettings)
	myLogGui.parent = p
	myLogGui.ready = true
	myLogGui.hidden = true
	myLogGui.groups = make(map[string]*flagGroup)
	return myLogGui
}

func (ls *LogSettings) Set(b bool) {
	// log.Set(ls.name, b)
	log.Warn("log.Set() FIXME: not working here anymore")
}

// Returns true if the status is valid
func (d *LogSettings) Ready() bool {
	if d == nil {return false}
	if ! d.parent.Ready() {return false}
	return d.ready
}

func (d *LogSettings) Update() bool {
	if ! d.Ready() {return false}
	return true
}

func (d *LogSettings) ShowFlags() {
	log.ShowFlags()
	return
}

func (d *LogSettings) SetAll(b bool) {
	log.SetAll(b)
	return
}
