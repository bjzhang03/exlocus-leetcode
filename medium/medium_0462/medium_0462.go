package medium_0462

import (
	"math"
	"sort"
)

func minMoves2(nums []int) int {
	if len(nums) > 0 {
		// 先排序
		sort.Ints(nums)
		save := make([]int, len(nums))
		sum := 0
		for i := 0; i < len(nums); i++ {
			sum = sum + nums[i]
			save[i] = sum
		}
		// 计算最小值
		result := math.MaxInt32
		for i := 0; i < len(nums); i++ {
			// 计算前面的差距
			left := 0
			if i > 0 {
				left = nums[i]*i - save[i-1]
			}
			// 计算后面的差距
			right := 0
			if i < len(nums) {
				right = save[len(nums)-1] - save[i] - (nums[i] * (len(nums) - 1 - i))

			}
			if result > (left + right) {
				result = (left + right)
			}
		}
		return result
	}
	return 0
}
