package medium_0456

import "math"

func find132pattern(nums []int) bool {
	if len(nums) >= 3 {
		curmin := math.MaxInt32
		save := make([][2]int, len(nums))
		for i := 0; i < len(nums); i++ {
			if curmin > nums[i] {
				curmin = nums[i]
			}
			save[i] = [2]int{curmin, nums[i]}
		}
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] > save[i][0] && nums[j] < save[i][1] {
					return true
				}
			}
		}
	}
	return false
}
