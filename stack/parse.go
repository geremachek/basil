package stack

import (
	"math"
	"strings"
	"strconv"
)

func parseValue(v string) (float64, error) {
	switch strings.ToLower(v) {
		case "e":   return math.E, nil
		case "pi":  return math.Pi, nil
		case "phi": return math.Phi, nil
		default:    return strconv.ParseFloat(v, 64)
	}
}

func (s *Stack) parseCommand(c Command) (err error) {
	switch c {
		case Pop:      _, err = s.pop()
		case Swap:     err = s.swap()
		case Recall:   err = s.recall()
		case Store:    err = s.store()
		case Clear:          s.clear()
		case Add:      err = s.operate(2, add)
		case Subtract: err = s.operate(2, subtract)
		case Multiply: err = s.operate(2, multiply)
		case Divide:   err = s.operate(2, divide)
	}

	return
}

func (s *Stack) Parse(input string) error {
	if v, e := parseValue(input); e == nil {
		s.push(v)
	} else if c, e := NewCommand(input); e == nil {
		s.parseCommand(c)
	} else {
		return e
	}

	return nil
}
