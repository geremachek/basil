package commands

import "github.com/geremachek/basil/berrs"

type Command int

const (
	Pop Command = iota
	Swap
	Recall
	Store
	Clear
	Angle
	Add
	Subtract
	Multiply
	Divide
	Mod
	Power
	Square
	Cube
	Sqrt
	Log
	Ln
	Logx
	Flip
	Sin
	Cos
	Tan
	Asin
	Acos
	Atan
)

// create a new command

func NewCommand(s string) (Command, error) {
	switch {
		case s == "pop"    || s == "p" || s == ".": return Pop, nil
		case s == "swap"   || s == "s" || s == ",": return Swap, nil
		case s == "recall" || s == "r" || s == "]": return Recall, nil
		case s == "store"  || s == "S" || s == "[": return Store, nil
		case s == "clear"  || s == "c" || s == ";": return Clear, nil
		case s == "angle"  || s == "a" || s == "<": return Angle, nil
		default: {
			switch s {
				case "+":      return Add, nil
				case "-":      return Subtract, nil
				case "*":      return Multiply, nil
				case "/":      return Divide, nil
				case "%":      return Mod, nil
				case "^":      return Power, nil
				case "square": return Square, nil
				case "cube":   return Cube, nil
				case "sqrt":   return Sqrt, nil
				case "log":    return Log, nil
				case "ln":     return Ln, nil
				case "logx":   return Logx, nil
				case "flip":   return Flip, nil
				case "sin":    return Sin, nil
				case "cos":    return Cos, nil
				case "tan":    return Tan, nil
				case "asin":   return Asin, nil
				case "acos":   return Acos, nil
				case "atan":   return Atan, nil
				default:       return Pop, berrs.ErrInvalidInput
			}
		}
	}
} 
