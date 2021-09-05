package stack

import (
	"math"
	"github.com/geremachek/basil/berrs"
)

// stack struct

type Stack struct {
	stack []float64

	memory float64
	Radians bool
}

// create a new stack

func NewStack() *Stack {
	return &Stack { nil, math.NaN(), true }
}

// check if the stack is empty

func (s Stack) Empty() bool {
	return len(s.stack) == 0
}

// get the size of the stack

func (s Stack) Size() int {
	return len(s.stack)
}

func (s Stack) Get(i int) float64 {
	return s.stack[i]
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

	// switch values

	var (
		l = len(s.stack)-1
		sl = l - 1
	)
	
	s.stack[l], s.stack[sl] = s.stack[sl], s.stack[l]

	return nil
}

// clear the stack

func (s *Stack) clear() {
	s.stack = nil
}

// change the angle mode

func (s *Stack) angle() {
	s.Radians = !s.Radians
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

	// pop off nums for the operation

	a, _ := s.pop()
	b, _ := s.pop()

	return s.runOperation(o(b, a), []float64{b, a})
}

// run and handle an operation on the stack

func (s *Stack) runOperation(v float64, args []float64) error {
	// if the operation failed, re-add the lost values

	if math.IsInf(math.Abs(v), 1) || math.IsNaN(v) {
		for i := 0; i < len(args); i++ {
			s.push(args[i])
		}

		return berrs.ErrArithmeticError
	}
	
	s.push(v)

	return nil
}

// deal with angle types...

func (s Stack) checkAngle(angle float64, toDeg bool) float64 {
	if s.Radians {
		return angle // input is already in radians
	}

	c := math.Pi / 180

	if toDeg { c = 180 / math.Pi }

	return angle * c // we need to convert it
}

func (s Stack) sin(angle float64) float64 { return math.Sin(s.checkAngle(angle, false)) }
func (s Stack) cos(angle float64) float64 { return math.Cos(s.checkAngle(angle, false)) }
func (s Stack) tan(angle float64) float64 { return math.Tan(s.checkAngle(angle, false)) }

func (s Stack) asin(v float64) float64 { return s.checkAngle(math.Asin(v), true) }
func (s Stack) acos(v float64) float64 { return s.checkAngle(math.Acos(v), true) }
func (s Stack) atan(v float64) float64 { return s.checkAngle(math.Atan(v), true) }
