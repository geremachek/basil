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
	Root
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
	switch s {
		case "pop",    "p", ".": return Pop, nil
		case "swap",   "s", ",": return Swap, nil
		case "recall", "r", "]": return Recall, nil
		case "store",  "S", "[": return Store, nil
		case "clear",  "c", ";": return Clear, nil
		case "angle",  "a", "<": return Angle, nil
		case "+":                return Add, nil
		case "-":                return Subtract, nil
		case "*":                return Multiply, nil
		case "/":                return Divide, nil
		case "%":                return Mod, nil
		case "^":                return Power, nil
		case "square":           return Square, nil
		case "cube":             return Cube, nil
		case "sqrt":             return Sqrt, nil
		case "log":              return Log, nil
		case "ln":               return Ln, nil
		case "logx":             return Logx, nil
		case "root":             return Root, nil
		case "flip":             return Flip, nil
		case "sin":              return Sin, nil
		case "cos":              return Cos, nil
		case "tan":              return Tan, nil
		case "asin":             return Asin, nil
		case "acos":             return Acos, nil
		case "atan":             return Atan, nil
		default:                 return Pop, berrs.ErrInvalidInput
	}
} 
