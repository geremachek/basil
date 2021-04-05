package stack

import (
	"math"
	"errors"
)

var (
	NAN = math.NaN()

	ErrTooFewArguments = errors.New("ErrTooFewArguments")
	ErrArithmeticError = errors.New("ErrArithmeticError")
	ErrInvalidInput    = errors.New("ErrInvalidInput")
	ErrNoStoredValue   = errors.New("ErrNoStoredValue")
)
