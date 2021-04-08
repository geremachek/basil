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

	if u, e := ui.NewUi(height); e == nil {
		u.Start()
	} else {
		fmt.Fprintf(os.Stderr, "basil: can't start interface\n%s\n", e)
	}
}
