package ui

import (
	"strings"
	"github.com/gdamore/tcell"
)

type LineBuff struct {
	buffer string
	
	locX int
	locY int
}

func NewLineBuff(x, y int) LineBuff {
	return LineBuff { "", x, y }
}

func (lb LineBuff) showPos(s tcell.Screen) {
	s.ShowCursor(lb.locX, lb.locY)
}

func (lb *LineBuff) Push(s tcell.Screen, ch rune) {
	lb.buffer += string(ch)
	s.SetContent(lb.locX, lb.locY, ch, []rune(""), tcell.StyleDefault)

	lb.locX++
	lb.showPos(s)
}

func (lb *LineBuff) Delete(s tcell.Screen) {
	if l := len(lb.buffer); l > 0 {
		lb.buffer = lb.buffer[:l-1]
		lb.locX--

		s.SetContent(lb.locX, lb.locY, ' ', []rune(""), tcell.StyleDefault)
		lb.showPos(s)
	}
}

func (lb *LineBuff) Refresh(s tcell.Screen) (toParse string) {
	lb.locX = 0

	lb.showPos(s)
	addstr(s, tcell.StyleDefault, 0, lb.locY, strings.Repeat(" ", len(lb.buffer))

	toParse = lb.buffer
	lb.buffer = ""

	return
}
