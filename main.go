package main

import (
	"github.com/geremachek/basil/stack"
)

func main() {
	s := stack.NewStack()
	s.Parse("15")
	s.Parse("5")
	s.Parse(";")

	s.D()
}
