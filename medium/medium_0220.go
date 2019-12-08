package medium

import (
	"math"
)

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	result := false
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j <= i+k && j < len(nums); j++ {
				if int(math.Abs(float64(nums[j]-nums[i]))) <= t {
					result = true
				}
			}
		}
	}
	return result
}
