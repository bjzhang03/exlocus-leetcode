package medium_0396

import "math"

func maxRotateFunction(A []int) int {
	result := 0
	if len(A) > 0 {
		// 取最小值的数据
		result = math.MinInt64
		doubleA := append([]int{}, A...)
		doubleA = append(doubleA, A...)
		for i := 0; i < len(A); i++ {
			sum := 0
			for j := i; j < i+len(A); j++ {
				sum = sum + doubleA[j]*(j-i)
			}
			if result < sum {
				result = sum
			}
		}
	}
	return result
}
