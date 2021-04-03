package stack

import (
	"errors"
)

// stack struct

type Stack struct {
	stack []float64
}

// create a new stack

func NewStack() Stack {
	return Stack { []float64{} }
}

// append some value to the stack

func (s *Stack) push(n float64) {
	s.stack = append(s.stack, n)
}

// remove and return the last element of the stack

func (s *Stack) pop() float64 {
	len := len(s.stack)-1
	last := s.stack[len]
	s.stack = s.stack[:len]

	return last
}

func (s *Stack) operate(req uint8, o func(...float64)) {
	if len(s.stack) < req {
		return errors.New("stack contains too few elements")
	}

	args := make([]float64, req)

	for i := 0; i < req; i++ {
		insert(args, s.pop(), 0)
	}

	s.push(o(args))

	return nil
}
