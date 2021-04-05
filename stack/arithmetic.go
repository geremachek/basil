package stack

import "math"

func add      (v []float64) float64 { return v[0] + v[1] }
func subtract (v []float64) float64 { return v[0] - v[1] }
func multiply (v []float64) float64 { return v[0] * v[1] }

func divide   (v []float64) float64 {
	if (v[1] == 0) {
		return NAN
	}

	return v[0] / v[1]
}

func sin      (v []float64) float64 { return math.Sin(v[0]) }
func tan      (v []float64) float64 { return math.Tan(v[0]) }
func cos      (v []float64) float64 { return math.Cos(v[0]) }

func asin     (v []float64) float64 { return math.Asin(v[0]) }
func atan     (v []float64) float64 { return math.Atan(v[0]) }
func acos     (v []float64) float64 { return math.Acos(v[0]) }
