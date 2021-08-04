package ui

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/gdamore/tcell"
)

const width = 20

var (
	strw = strconv.Itoa(width)
	bar = strings.Repeat("─", width) // barrier decoration string
)

// draw text to the screen

func addstr(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	arr := []rune(text)

	for i := x; i < len(arr)+x; i++ {
		s.SetContent(i, y, arr[i-x], []rune(""), style)
	}
}

// draw the barrier decoration...

func drawLine(s tcell.Screen, x, y int) {
	addstr(s, tcell.StyleDefault, x, y, bar)
}

// allign and trim our text

func drawAligned(s tcell.Screen, x, y int, text string) {
	disp := text

	// if there isn't enough room, add "..."

	if l := len(disp); l > width {
		disp = disp[:l - (l - width) - 3] + "..."
	}

	addstr(s, tcell.StyleDefault, x, y, fmt.Sprintf("%" + strw + "s", disp))
}
