package main

import (
	"os"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/geremachek/basil/ui"
)

func main() {
	var height int

	pflag.IntVarP(&height, "height", "H", 10, "set how many elements are displayed at once")
	pflag.Parse()

	// start the ui, trapping errors

	if u, e := ui.NewUi(height); e == nil {
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
