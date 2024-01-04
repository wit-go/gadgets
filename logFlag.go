/*
	A log.Flag

	-----------------------------------------------
	|            |                                 |
	|   [ X ]    |    INFO (controls log.Info()    |
	|            |                                 |
	-----------------------------------------------
*/
package gadgets

import 	(
	"go.wit.com/log"
	"go.wit.com/gui/gui"
)

type LogFlag struct {
	p	*gui.Node	// parent widget
	c	*gui.Node	// checkbox widget

	name string
	subsystem string
	desc string
	b bool

	Custom func()
}

func (f *LogFlag) Get() bool {
	return log.Get(f.subsystem, f.name)
}

func (f *LogFlag) Set(b bool) {
	log.Println("LogFlag.Set() =", b)
	log.Set(f.subsystem, f.name, b)
}

func NewLogFlag(n *gui.Node, lf *log.LogFlag) *LogFlag {
	f := LogFlag {
		name: lf.Name,
		subsystem: lf.Subsystem,
		desc: lf.Desc,
		p: n,
	}

	// various timeout settings
	f.c = n.NewCheckbox(f.name + " (" + f.desc + ")")
	f.c.Custom = func() {
		log.Set(f.subsystem, f.name, f.c.B)
		log.Println("LogFlag.Custom() user changed value to =", log.Get(f.subsystem, f.name))
	}

	return &f
}
