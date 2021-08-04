package arithmetic

import "math"

// Invalid number

var NaN = math.NaN()

func Add      (a, b float64) float64 { return a + b }
func Subtract (a, b float64) float64 { return a - b }
func Multiply (a, b float64) float64 { return a * b }

func Divide   (a, b float64) float64 {
	// handle division by zero!

	if (b == 0) {
		return NaN
	}

	return a / b
}

func Square   (v float64) float64    { return math.Pow(v, 2) }
func Cube     (v float64) float64    { return math.Pow(v, 3) }
func Logx     (a, b float64) float64 { return math.Log10(b) / math.Log10(a) }
func Root     (a, b float64) float64 { return math.Pow(b, 1 / a) }
func Flip     (v float64) float64    { return -v }

func FactW    (v float64) float64 {
	// no bignum for now

	if v < 21 {
		return float64(fact(uint64(v)))
	}

	return NaN
}

func fact     (v uint64) uint64 {
	if v == 0 || v == 1 { // base case
		return 1
	}

	return v * fact(v-1)
}
