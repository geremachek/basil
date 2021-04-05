package berrs

import "errors"

var (
	ErrTooFewArguments = errors.New("ErrTooFewArguments")
	ErrArithmeticError = errors.New("ErrArithmeticError")
	ErrInvalidInput    = errors.New("ErrInvalidInput")
	ErrNoStoredValue   = errors.New("ErrNoStoredValue")
)
