package main

import (
	"os"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/geremachek/basil/ui"
)

func main() {
	var (
		height int
		width int
	)

	pflag.IntVarP(&height, "height", "H", 10, "set how many elements are displayed at once")
	pflag.IntVarP(&width, "width", "w", 30, "set the width of the stack window")
	pflag.Parse()

	// start the ui, trapping errors

	if u, e := ui.NewUi(height, width); e == nil {
		if err := u.Start(); err != nil {
			printErr(e)
		}
	} else {
		printErr(e)
	}
}

func printErr(e error) {
	fmt.Fprintf(os.Stderr, "basil: can't start interface\n%s\n", e)
	os.Exit(1)
}
