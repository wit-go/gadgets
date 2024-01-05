package logsettings

import 	(
	"go.wit.com/gui/gui"
	"go.wit.com/gui/gadgets"
)

var myLogGui *LogSettings

type LogSettings struct {
	ready	bool
	hidden	bool
	err	error

	groups map[string]*flagGroup

	parent	*gui.Node // where to draw our window
	win	*gadgets.BasicWindow // our window for displaying the log package settings

	buttonG	*gui.Node // the group of buttons
	flagG	*gui.Node // the group of all the flag checkbox widgets
}
