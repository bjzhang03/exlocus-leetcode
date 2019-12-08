package medium

import (
	"math"
)

func myPow(x float64, n int) float64 {
	result := 0.0
	if x != 0 {
		if n == 0 {
			result = 1
		} else {
			result = math.Pow(x, float64(n))
		}
	}
	return result
}
