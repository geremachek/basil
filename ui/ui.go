package ui

import (
	"strings"
	"strconv"
	"github.com/gdamore/tcell"
	"github.com/geremachek/basil/stack"
)

var clearLine = strings.Repeat(" ", width)

type ui struct {
	stack *stack.Stack
	buff *lineBuff

	scr tcell.Screen
	height int
	running bool
}

// create a new ui struct

func NewUi(h int) (*ui, error) {
	if s, err := tcell.NewScreen(); err == nil {
		return &ui { stack.NewStack(), newLineBuff(0, h+2), s, h, true }, nil
	} else {
		return &ui{}, err
	}
}

// draw the window showing the stack

func (u *ui) drawStackWindow() {
	if !u.stack.Empty() {
		y := u.height-1

		for i := u.stack.Size()-1; i >= 0 && y >= 0; i-- {
			drawAligned(u.scr, 0, y, strconv.FormatFloat(u.stack.Get(i), 'f', -1, 64))
			y--
		}
	}
}

// clear the stack window

func (u *ui) clearStackWindow(lines int) {
	for y := u.height-1; y >= u.height-lines; y-- {
		addstr(u.scr, tcell.StyleDefault, 0, y, clearLine)
	}
}

// draw the angle decoration

func (u ui) drawAngleMode() {
	m := 'D'

	// R is displayed when in radian mode

	if u.stack.Radians {
		m = 'R'
	}

	u.scr.SetContent(width+1, u.height, m, []rune(""), tcell.StyleDefault)
}

// match keys with actions

func (u *ui) matchKeys(input *tcell.EventKey) {
	switch input.Key() {
		case tcell.KeyESC:                     u.running = false
		case tcell.KeyEnter:                   u.parseLine()
		case tcell.KeyDEL, tcell.KeyBackspace: u.buff.delete(u.scr)
		case tcell.KeyRune:                    u.buff.push(u.scr, input.Rune())
	}

	// only refresh the screen if the interface is still "running"

	if u.running {
		u.scr.Show()
	}
}

// pars a line of input, redrawing the screen

func (u *ui) parseLine() {
	before := u.stack.Size() // old stack size

	if e := u.stack.Parse(u.buff.text()); e == nil { // redraw if parsing is successful
		u.clearStackWindow(before)
		u.drawStackWindow()
		u.drawAngleMode()

		u.buff.refresh(u.scr)
	}
}

// draw the screen and listen for input

func (u *ui) Start() error {
	if e := u.scr.Init(); e == nil {
		// draw to the screen
		
		drawLine(u.scr, 0, u.height)
		u.drawAngleMode()
		
		u.scr.ShowCursor(0, u.height + 2)
		u.scr.Show()

		// handle input

		var input tcell.Event

		for u.running {
			input = u.scr.PollEvent()

			switch k := input.(type) {
				case *tcell.EventKey: u.matchKeys(k)
			}
		}

		u.scr.Fini()
	} else {
		return e
	}

	return nil
}
