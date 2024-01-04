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

	Name string
	Subsystem string
	Desc string
	b bool

	Custom func()
}

func (f *LogFlag) Get() bool {
	return log.Get(f.Subsystem, f.Name)
}

func (f *LogFlag) Set(b bool) {
	log.Println("LogFlag.Set() =", b)
	log.Set(f.Subsystem, f.Name, b)
}

func NewLogFlag(n *gui.Node, lf *log.LogFlag) *LogFlag {
	f := LogFlag {
		Name: lf.Name,
		Subsystem: lf.Subsystem,
		Desc: lf.Desc,
		p: n,
	}

	// various timeout settings
	f.c = n.NewCheckbox(f.Name + ": " + f.Desc)
	f.c.Custom = func() {
		log.Set(f.Subsystem, f.Name, f.c.B)
		log.Println("LogFlag.Custom() user changed value to =", log.Get(f.Subsystem, f.Name))
	}

	return &f
}
