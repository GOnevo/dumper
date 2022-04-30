package dumper

import (
	"github.com/sanity-io/litter"
	"os"
)

// D prints pretty dump to stdout only.
func D(v ...interface{}) {
	dump(false, v...)
}

// DD prints pretty dump to stdout and die.
func DD(v ...interface{}) {
	dump(true, v...)
}

func dump(exit bool, v ...interface{}) {
	litter.Dump(v...)
	osExit(exit)
}

func osExit(yes bool) {
	if yes {
		os.Exit(0)
	}
}
