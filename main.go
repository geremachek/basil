package main

import (
	"os"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/geremachek/basil/ui"
)

func main() {
	var height int

	pflag.IntVarP(&height, "height", "H", 5, "set how many elements are displayed at once")
	pflag.Parse()

	message := "basil: can't start interface\n%s\n"

	// start the ui, trapping errors

	if u, e := ui.NewUi(height); e == nil {
		if err := u.Start(); err != nil {
			fmt.Fprintf(os.Stderr, message, e)
		}
	} else {
		fmt.Fprintf(os.Stderr, message, e)
	}
}
