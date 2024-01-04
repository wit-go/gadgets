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
	Default bool
	b bool

	Custom func()
}

func (f *LogFlag) Get() bool {
	return log.Get(f.Subsystem, f.Name)
}

func (f *LogFlag) Set(b bool) {
	log.Info("LogFlag.Set() =", b)
	log.Set(f.Subsystem, f.Name, b)
	f.c.Set(b)
}

func (f *LogFlag) SetDefault() {
	log.Info("LogFlag.SetDefault() =", f.Default)
	log.Set(f.Subsystem, f.Name, f.Default)
	f.c.Set(f.Default)
}

func NewLogFlag(n *gui.Node, lf *log.LogFlag) *LogFlag {
	f := LogFlag {
		Name: lf.Name,
		Subsystem: lf.Subsystem,
		Desc: lf.Desc,
		Default: lf.Default,
		p: n,
	}

	// various timeout settings
	f.c = n.NewCheckbox(f.Name + ": " + f.Desc)
	f.c.Custom = func() {
		log.Set(f.Subsystem, f.Name, f.c.B)
		log.Info("LogFlag.Custom() user changed value to =", log.Get(f.Subsystem, f.Name))
	}
	f.c.Set(lf.B)

	return &f
}
