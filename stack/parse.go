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
		case cmd.Angle:          s.angle()
		case cmd.Add:      err = s.operateDouble(ari.Add)
		case cmd.Subtract: err = s.operateDouble(ari.Subtract)
		case cmd.Multiply: err = s.operateDouble(ari.Multiply)
		case cmd.Divide:   err = s.operateDouble(ari.Divide)
		case cmd.Mod:      err = s.operateDouble(math.Mod)
		case cmd.Power:    err = s.operateDouble(math.Pow)
		case cmd.Square:   err = s.operateSingle(ari.Square)
		case cmd.Cube:     err = s.operateSingle(ari.Cube)
		case cmd.Sqrt:     err = s.operateSingle(math.Sqrt)
		case cmd.Log:      err = s.operateSingle(math.Log10)
		case cmd.Ln:       err = s.operateSingle(math.Log)
		case cmd.Logx:     err = s.operateDouble(ari.Logx)
		case cmd.Flip:     err = s.operateSingle(ari.Flip)
		case cmd.Sin:      err = s.operateSingle(s.sin)
		case cmd.Cos:      err = s.operateSingle(s.cos)
		case cmd.Tan:      err = s.operateSingle(s.tan)
		case cmd.Asin:     err = s.operateSingle(s.asin)
		case cmd.Acos:     err = s.operateSingle(s.acos)
		case cmd.Atan:     err = s.operateSingle(s.atan)
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
