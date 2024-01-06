package gadgets

// initializes logging and command line options

import (
	"go.wit.com/log"
)

var INFO log.LogFlag

func init() {
	INFO.B = false
	INFO.Name = "INFO"
	INFO.Subsystem = "gadgets"
	INFO.Short = "gadgets"
	INFO.Desc = "general info"
	INFO.Register()
}
