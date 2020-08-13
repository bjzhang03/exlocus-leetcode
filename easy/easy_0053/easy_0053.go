package easy_0053

import (
	"math"
)

func maxSubArray(nums []int) int {
	result := 0
	max := math.MinInt64
	for i := 0; i < len(nums); i++ {
		result = result + nums[i]
		if max < result {
			max = result
		}
		if result < 0 {
			result = 0
		}
	}
	return max
}
