package medium_0611

import "sort"

func triangleNumber(nums []int) int {
	if len(nums) >= 3 {
		result := 0
		sort.Ints(nums)
		for i := 0; i < len(nums); i++ {
			for j := i + 1; j < len(nums); j++ {
				for k := j + 1; k < len(nums); k++ {
					if isTriangle(nums[i], nums[j], nums[k]) {
						result++
					}
					// 优化处理
					if nums[k] >= nums[i]+nums[j] {
						break
					}
				}
			}
		}
		return result
	}
	return 0
}

func isTriangle(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	}
	return false
}
