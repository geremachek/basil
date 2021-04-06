ipackage stack

import (
	"math"
	"fmt"
	"github.com/geremachek/basil/berrs"
)

// stack struct

type Stack struct {
	stack []float64

	memory float64
	radians bool
}

// create a new stack

func NewStack() Stack {
	return Stack { []float64{}, NAN, false }
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

	return 0.0, berrs.ErrTooFewArguments
}

// swap two values

func (s *Stack) swap() error {
	if len(s.stack) < 2 {
		return berrs.ErrTooFewArguments
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
		return berrs.ErrNoStoredValue
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

// single element operation

func (s *Stack) operateSingle(o func(float64) float64) error {
	if v, e := s.pop(); e == nil {
		return s.runOperation(o(v), []float64{v})
	} else {
		return e
	}
}

// two element operation

func (s *Stack) operateDouble(o func(float64, float64) float64) error {
	if len(s.stack) < 2 {
		return berrs.ErrTooFewArguments // operate did nothing
	}

	a, _ := s.pop()
	b, _ := s.pop()

	return s.runOperation(o(b, a), []float64{b, a})
}

// run and handle an operation on the stack

func (s *Stack) runOperation(v float64, args []float64) error {
	if math.IsInf(math.Abs(v), 1) || math.IsNaN(v) {
		for i := 0; i < len(args); i++ {
			s.push(args[i])
		}

		return berrs.ErrArithmeticError
	}
	
	s.push(v)

	return nil
}

func (s Stack) checkAngle(angle float64, toRad bool) {
	if s.radians {
		return angle // input is already in radians
	}

	c := math.Pi / 180

	if toRad { c = 180 / math.Pi }

	return angle * c // we need to convert it
}

func (s Stack) sin(angle float64) float64 { return math.Sin(checkAngle(angle, false)) }
func (s Stack) cos(angle float64) float64 { return math.Cos(checkAngle(angle, false)) }
func (s Stack) tan(angle float64) float64 { return math.Tan(checkAngle(angle, false)) }
