package ui

import (
	"fmt"
	"strings"
	"github.com/gdamore/tcell"
)

const WIDTH = 20

// draw text to the screen

func addstr(s tcell.Screen, style tcell.Style, x int, y int, text string) {
	arr := []rune(text)

	for i := x; i < len(arr)+x; i++ {
		s.SetContent(i, y, arr[i-x], []rune(""), style)
	}
}

func drawLine(s tcell.Screen, x, y int) {
	addstr(s, tcell.StyleDefault, x, y, strings.Repeat("~", WIDTH))
}

func drawAligned(s tcell.Screen, x, y int, text string) {
	disp := text

	if l := len(disp); l > WIDTH {
		disp = disp[:l - (l - WIDTH) - 3] + "..."
	}

	addstr(s, tcell.StyleDefault, x, y, fmt.Sprintf("%20s", disp))
}
