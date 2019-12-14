package medium_0029

import (
	"math"
)

func preHandle(dividend int, divisor int) (bool, int, int) {
	result := false
	if dividend > 0 && divisor > 0 {
		result = true
	} else if dividend < 0 && divisor < 0 {
		result = true
	}
	return result, int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
}

func xorAdd(dividend int64, divisor int64) int {
	if dividend < divisor {
		return 0
	}
	
	result := 1
	tmpDivisor := divisor
	for tmpDivisor <= dividend {
		tmpDivisor = tmpDivisor << 1
		result = result << 1
	}
	result = result >> 1

	return result + xorAdd(dividend-(tmpDivisor>>1), divisor)
}

func divide(dividend int, divisor int) int {
	result := 0
	negativeFlag, dividend, divisor := preHandle(dividend, divisor)

	if negativeFlag {
		result = xorAdd(int64(dividend), int64(divisor))
		if result >= math.MaxInt32 {
			result = math.MaxInt32
		}
	} else {
		result = -xorAdd(int64(dividend), int64(divisor))
		if result <= math.MinInt32 {
			result = math.MinInt32
		}
	}

	return result
}
