package ui

import (
	"strings"
	"github.com/gdamore/tcell"
)

type lineBuff struct {
	buffer string
	
	locX int
	locY int
}

func newLineBuff(x, y int) lineBuff {
	return lineBuff { "", x, y }
}


func (lb lineBuff) text() string {
	return lb.buffer
}

// show our cursor

func (lb lineBuff) showPos(s tcell.Screen) {
	s.ShowCursor(lb.locX, lb.locY)
}

// add a char...

func (lb *lineBuff) push(s tcell.Screen, ch rune) {
	lb.buffer += string(ch)
	s.SetContent(lb.locX, lb.locY, ch, []rune(""), tcell.StyleDefault)

	lb.locX++
	lb.showPos(s)
}

// delete our char

func (lb *lineBuff) delete(s tcell.Screen) {
	if l := len(lb.buffer); l > 0 {
		lb.buffer = lb.buffer[:l-1]
		lb.locX--

		s.SetContent(lb.locX, lb.locY, ' ', []rune(""), tcell.StyleDefault)
		lb.showPos(s)
	}
}

// reset the line

func (lb *lineBuff) refresh(s tcell.Screen) {
	addstr(s, tcell.StyleDefault, 0, lb.locY, strings.Repeat(" ", len(lb.buffer)))

	lb.locX = 0
	lb.buffer = ""
	
	lb.showPos(s)
}
