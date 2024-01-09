package gadgets

// initializes logging and command line options

import (
	"go.wit.com/log"
)

var INFO *log.LogFlag

func init() {
	full := "go.wit.com/gui/gadget"
	short := "gadgets"

	INFO = log.NewFlag("INFO", false, full, short, "General Info")
}
