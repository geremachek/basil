package ui

import (
	"fmt"
	"strconv"
	"github.com/gdamore/tcell"
)

// draw text to the screen

func addstr(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	arr := []rune(text)

	for i := x; i < len(arr)+x; i++ {
		s.SetContent(i, y, arr[i-x], []rune(""), style)
	}
}

// allign and trim our text

func drawAligned(s tcell.Screen, x, y, width int, text string) {
	disp := text

	// if there isn't enough room, add "..."

	if l := len(disp); l > width {
		disp = disp[:l - (l - width) - 3] + "..."
	}

	addstr(s, tcell.StyleDefault, x, y, fmt.Sprintf("%" + strconv.Itoa(width) + "s", disp))
}
