package stack

type Command int

const (
	Pop Command = iota
	Swap
	Recall
	Store
	Clear
	Add
	Subtract
	Multiply
	Divide
)

// create a new command

func NewCommand(s string) (Command, error) {
	switch {
		case s == "pop"    || s == "p" || s == ".": return Pop, nil
		case s == "swap"   || s == "s" || s == ",": return Swap, nil
		case s == "recall" || s == "r" || s == "]": return Recall, nil
		case s == "store"  || s == "S" || s == "[": return Store, nil
		case s == "clear"  || s == "c" || s == ";": return Clear, nil
		default: {
			switch s {
				case "+": return Add, nil
				case "-": return Subtract, nil
				case "*": return Multiply, nil
				case "/": return Divide, nil
				default:  return Pop, ErrInvalidInput
			}
		}
	}
} 
