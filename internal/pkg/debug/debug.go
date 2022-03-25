package debug

import "log"

const Debug = true

func DPrinf(format string, a ...interface{}) {
	if Debug {
		log.Printf(format, a...)
	}
}
