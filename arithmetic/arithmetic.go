package arithmetic

import "math"

func Add      (a, b float64) float64 { return a + b }
func Subtract (a, b float64) float64 { return a - b }
func Multiply (a, b float64) float64 { return a * b }

func Divide   (a, b float64) float64 {
	if (b == 0) {
		return math.NaN()
	}

	return a / b
}

func Square   (v float64) float64    { return math.Pow(v, 2) }
func Cube     (v float64) float64    { return math.Pow(v, 3) }
func Logx     (a, b float64) float64 { return math.Log10(b) / math.Log10(a) }
func Root     (a, b float64) float64 { return math.Pow(b, 1 / a) }
func Flip     (v float64) float64    { return -v }

func FactW    (v float64) float64    { return float64(Fact(int(math.Trunc(v)))) }

func Fact     (v int) int {
	if v == 0 || v == 1 {
		return v
	}

	return v * Fact(v-1)
}
