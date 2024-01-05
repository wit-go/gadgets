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
	myLogGui.groups = make(map[string]*flagGroup)
	myLogGui.ready = true
	myLogGui.hidden = true
	return myLogGui
}

// Returns true if the status is valid
func (d *LogSettings) Ready() bool {
	if d == nil {return false}
	if ! d.parent.Ready() {return false}
	if (d.win == nil) {
		d.draw()
	}
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
