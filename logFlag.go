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
	lf	*log.LogFlag

	Name string
	Subsystem string
	Desc string
	Default bool
	b bool

	Custom func()
}

func (f *LogFlag) Get() bool {
	return f.lf.Get()
}

func (f *LogFlag) Set(b bool) {
	log.Info("LogFlag.Set() =", b)
	f.lf.Set(b)
	f.c.Set(b)
}

func (f *LogFlag) SetDefault() {
	log.Info("LogFlag.SetDefault() =", f.Default)
	f.lf.SetDefault()
	f.c.Set(f.lf.Get())
}

func NewLogFlag(n *gui.Node, lf *log.LogFlag) *LogFlag {
	f := LogFlag {
		p: n,
	}
	f.Name = lf.GetName()
	f.Subsystem = lf.GetSubsystem()
	f.Desc = lf.GetDesc()

	// various timeout settings
	f.c = n.NewCheckbox(f.Name + ": " + f.Desc)
	f.c.Custom = func() {
		log.Set(f.Subsystem, f.Name, f.c.B)
		log.Info("LogFlag.Custom() user changed value to =", log.Get(f.Subsystem, f.Name))
	}
	f.c.Set(lf.Get())

	return &f
}
