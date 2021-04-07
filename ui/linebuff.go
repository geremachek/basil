package ui

import "github.com/gdamore/tcell"

type LineBuff struct {
	buffer string
	
	locX int
	locY int
}

func NewLineBuff(x, y int) LineBuff {
	return LineBuff { "", x, y }
}

func (lb *LineBuff) Push(s tcell.Screen, ch rune) {
	lb.buffer += string(ch)
	s.SetContent(lb.locX, lb.locY, ch, []rune(""), tcell.StyleDefault)

	lb.locX++
	s.ShowCursor(lb.locX, lb.locY)
}


