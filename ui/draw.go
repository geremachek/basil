package ui

import (
	"fmt"
	"strconv"
	"github.com/gdamore/tcell"
)

// draw text to the screen

func addString(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	curs := x

	for _, ch := range text {
		s.SetContent(curs, y, ch, []rune{}, style)
		curs++
	}
}

// allign and trim our text

func drawAligned(s tcell.Screen, x, y, width int, text string) {
	disp := text

	// if there isn't enough room, add "..."

	if l := len(disp); l > width {
		disp = disp[:l - (l - width) - 3] + "..."
	}

	addString(s, tcell.StyleDefault, x, y, fmt.Sprintf("%" + strconv.Itoa(width) + "s", disp))
}
