package stack

import (
	"math"
	"strings"
	"strconv"
	cmd "github.com/geremachek/basil/commands"
	ari "github.com/geremachek/basil/arithmetic"
)

func parseValue(v string) (float64, error) {
	switch strings.ToLower(v) {
		case "e":   return math.E, nil
		case "pi":  return math.Pi, nil
		case "phi": return math.Phi, nil
		default:    return strconv.ParseFloat(v, 64)
	}
}

func (s *Stack) parseCommand(c cmd.Command) (err error) {
	switch c {
		case cmd.Pop:      _, err = s.pop()
		case cmd.Swap:     err = s.swap()
		case cmd.Recall:   err = s.recall()
		case cmd.Store:    err = s.store()
		case cmd.Clear:          s.clear()
		case cmd.Add:      err = s.operateDouble(ari.Add)
		case cmd.Subtract: err = s.operateDouble(ari.Subtract)
		case cmd.Multiply: err = s.operateDouble(ari.Multiply)
		case cmd.Divide:   err = s.operateDouble(ari.Divide)
		case cmd.Mod:      err = s.operateDouble(math.Mod)
		case cmd.Square:   err = s.operateSingle(ari.Square)
		case cmd.Cube:     err = s.operateSingle(ari.Cube)
	}

	return
}

func (s *Stack) Parse(input string) error {
	if v, e := parseValue(input); e == nil {
		s.push(v)
	} else if c, e := cmd.NewCommand(input); e == nil {
		s.parseCommand(c)
	} else {
		return e
	}

	return nil
}
