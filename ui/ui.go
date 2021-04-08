package ui

import (
	"strings"
	"strconv"
	"github.com/gdamore/tcell"
	"github.com/geremachek/basil/stack"
)

type Ui struct {
	stack stack.Stack
	buff LineBuff

	scr tcell.Screen
	height int
	running bool
}

func NewUi(h int) (Ui, error) {
	if s, err := tcell.NewScreen(); err == nil {
		return Ui { stack.NewStack(), NewLineBuff(0, h+2), s, h, true }, nil
	} else {
		return Ui{}, err
	}
}

func (u *Ui) drawStackWindow() {
	if !u.stack.Empty() {
		y := u.height-1

		for i := u.stack.Size()-1; i >= 0 && y >= 0; i-- {
			drawAligned(u.scr, 0, y, strconv.FormatFloat(u.stack.Get(i), 'f', -1, 64))
			y--
		}
	}
}

func (u *Ui) clearStackWindow(lines int) {
	spaces := strings.Repeat(" ", WIDTH)

	for y := u.height-1; y >= u.height-lines; y-- {
		addstr(u.scr, tcell.StyleDefault, 0, y, spaces)
	}
}

func (u Ui) drawAngleMode() {
	m := 'D'

	if u.stack.Radians {
		m = 'R'
	}

	u.scr.SetContent(WIDTH+1, u.height, m, []rune(""), tcell.StyleDefault)
}

func (u *Ui) matchKeys(input *tcell.EventKey) {
	switch input.Key() {
		case tcell.KeyCtrlLeftSq:              u.running = false
		case tcell.KeyEnter:                   u.parseLine()
		case tcell.KeyDEL, tcell.KeyBackspace: u.buff.Delete(u.scr)
		case tcell.KeyRune:                    u.buff.Push(u.scr, input.Rune())
	}

	u.scr.Show()
}

func (u *Ui) parseLine() {
	before := u.stack.Size()

	if e := u.stack.Parse(u.buff.Buffer()); e == nil {
		u.clearStackWindow(before)
		u.drawStackWindow()
		u.drawAngleMode()

		u.buff.Refresh(u.scr)
	}
}

func (u *Ui) Start() error {
	if e := u.scr.Init(); e == nil {
		drawLine(u.scr, 0, u.height)
		u.drawAngleMode()
		
		u.scr.ShowCursor(0, u.height + 2)
		u.scr.Show()

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
