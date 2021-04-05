package stack

import (
	"math"
	"fmt"
)

// stack struct

type Stack struct {
	stack []float64
	memory float64
}

// create a new stack

func NewStack() Stack {
	return Stack { []float64{}, NAN }
}

func (s *Stack) D() {
	fmt.Println(s.stack)
}

// append some value to the stack

func (s *Stack) push(n float64) {
	s.stack = append(s.stack, n)
}

// remove and return the last element of the stack

func (s *Stack) pop() (float64, error) {
	if len := len(s.stack)-1; len > -1 {
		last := s.stack[len]
		s.stack = s.stack[:len]

		return last, nil
	}

	return 0.0, ErrTooFewArguments
}

// swap two values

func (s *Stack) swap() error {
	if len(s.stack) < 2 {
		return ErrTooFewArguments
	}

	a, _ := s.pop()
	b, _ := s.pop()

	s.push(a)
	s.push(b)

	return nil
}

// clear the stack

func (s *Stack) clear() {
	s.stack = []float64{}
}

// recall the stored value

func (s *Stack) recall() error {
	if math.IsNaN(s.memory) {
		return ErrNoStoredValue
	}

	s.push(s.memory)

	return nil
}

// set the stored memory value

func (s *Stack) store() error {
	if v, e := s.pop(); e == nil {
		s.memory = v
	} else {
		return e
	}

	return nil
}

// operate on the stack

func (s *Stack) operate(req int, o func([]float64) float64) error {
	if len(s.stack) < req {
		return ErrTooFewArguments // operate did nothing
	}

	args := make([]float64, req)

	for i := 0; i < req; i++ {
		args = append(args, 0)
		copy(args[1:], args[0:])
		args[0], _ = s.pop()

	}

	if v := o(args); math.IsInf(math.Abs(v), 1)  || math.IsNaN(v) {
		for i := 0; i < req; i++ {
			s.push(args[i])
		}

		return ErrArithmeticError
	} else {
		s.push(v)
	}

	return nil
}
