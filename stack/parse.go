package stack

import (
	"math"
	"strings"
	"strconv"
	cmd "github.com/geremachek/basil/commands"
	ari "github.com/geremachek/basil/arithmetic"
)

// parse variables...

func parseValue(v string) (float64, error) {
	switch strings.ToLower(v) {
		case "e":   return math.E, nil
		case "pi":  return math.Pi, nil
		case "phi": return math.Phi, nil
		default:    return strconv.ParseFloat(v, 64) // should be a normal value, if not, an error!
	}
}

// parse commands...

func (s *Stack) parseCommand(c cmd.Command) {
	switch c {
		case cmd.Pop:      s.pop()
		case cmd.Swap:     s.swap()
		case cmd.Recall:   s.recall()
		case cmd.Store:    s.store()
		case cmd.Clear:    s.clear()
		case cmd.Angle:    s.angle()
		case cmd.Add:      s.operateDouble(ari.Add)
		case cmd.Subtract: s.operateDouble(ari.Subtract)
		case cmd.Multiply: s.operateDouble(ari.Multiply)
		case cmd.Divide:   s.operateDouble(ari.Divide)
		case cmd.Mod:      s.operateDouble(math.Mod)
		case cmd.Power:    s.operateDouble(math.Pow)
		case cmd.Square:   s.operateSingle(ari.Square)
		case cmd.Cube:     s.operateSingle(ari.Cube)
		case cmd.Sqrt:     s.operateSingle(math.Sqrt)
		case cmd.Log:      s.operateSingle(math.Log10)
		case cmd.Ln:       s.operateSingle(math.Log)
		case cmd.Logb:     s.operateDouble(ari.Logb)
		case cmd.Root:     s.operateDouble(ari.Root)
		case cmd.Flip:     s.operateSingle(ari.Flip)
		case cmd.Sin:      s.operateSingle(s.sin)
		case cmd.Cos:      s.operateSingle(s.cos)
		case cmd.Tan:      s.operateSingle(s.tan)
		case cmd.Asin:     s.operateSingle(s.asin)
		case cmd.Acos:     s.operateSingle(s.acos)
		case cmd.Atan:     s.operateSingle(s.atan)
		case cmd.Fact:     s.operateSingle(ari.FactW)
	}
}

// value... command... error? Deal with them all!

func (s *Stack) Parse(input string) {
	for _, atom := range strings.Fields(input) {
		if v, e := parseValue(atom); e == nil {
			s.push(v)
		} else if c, e := cmd.NewCommand(atom); e == nil {
			s.parseCommand(c)
		}
	}
}
