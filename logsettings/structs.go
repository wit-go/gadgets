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
	window	*gui.Node // our window for displaying the log package settings

	// Primary Directives
	status	*gadgets.OneLiner
	summary	*gadgets.OneLiner

	checkbox *gadgets.LogFlag
}
