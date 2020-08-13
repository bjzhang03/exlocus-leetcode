package medium_0679

import (
	"math"
)

func judgePoint24(nums []int) bool {
	if len(nums) == 4 {
		all := arranges(nums)
		for i := 0; i < len(all); i++ {
			if deepFirstSearch(tofloat64(all[i])) {
				return true
			}
		}
	}
	return false
}

// 深度优先遍历,测试当前的序列是否可以
func deepFirstSearch(nums []float64) bool {
	if len(nums) == 1 && math.Abs(nums[0]-float64(24)) < 0.0000001 {
		// 这里使用浮点数计算,存在精度损失的问题
		// 但是对付当前的题目,使用这个精度已经可以满足需求了
		// 算是我在这里偷懒了
		return true
	} else if len(nums) > 1 {
		for i := 0; i < len(nums) && i+1 < len(nums); i++ {
			first := nums[i]
			second := nums[i+1]
			if deepFirstSearch(append(append(append([]float64{}, nums[:i]...), first+second), nums[i+2:]...)) {
				return true
			}
			if deepFirstSearch(append(append(append([]float64{}, nums[:i]...), first-second), nums[i+2:]...)) {
				return true
			}
			if deepFirstSearch(append(append(append([]float64{}, nums[:i]...), first*second), nums[i+2:]...)) {
				return true
			}
			if second != 0 && deepFirstSearch(append(append(append([]float64{}, nums[:i]...), first/second), nums[i+2:]...)) {
				return true
			}
		}
	}
	return false
}

func tofloat64(nums []int) []float64 {
	return []float64{float64(nums[0]), float64(nums[1]), float64(nums[2]), float64(nums[3])}
}

// 获取所有的组合的数据
func arranges(nums []int) [][]int {
	result := [][]int{}
	if len(nums) == 4 {
		result = append(result, []int{nums[0], nums[1], nums[2], nums[3]})
		result = append(result, []int{nums[0], nums[1], nums[3], nums[2]})
		result = append(result, []int{nums[0], nums[2], nums[1], nums[3]})
		result = append(result, []int{nums[0], nums[2], nums[3], nums[1]})
		result = append(result, []int{nums[0], nums[3], nums[1], nums[2]})
		result = append(result, []int{nums[0], nums[3], nums[2], nums[1]})
		result = append(result, []int{nums[1], nums[0], nums[2], nums[3]})
		result = append(result, []int{nums[1], nums[0], nums[3], nums[2]})
		result = append(result, []int{nums[1], nums[2], nums[0], nums[3]})
		result = append(result, []int{nums[1], nums[2], nums[3], nums[0]})
		result = append(result, []int{nums[1], nums[3], nums[0], nums[2]})
		result = append(result, []int{nums[1], nums[3], nums[2], nums[0]})
		result = append(result, []int{nums[2], nums[0], nums[1], nums[3]})
		result = append(result, []int{nums[2], nums[0], nums[3], nums[1]})
		result = append(result, []int{nums[2], nums[1], nums[0], nums[3]})
		result = append(result, []int{nums[2], nums[1], nums[3], nums[0]})
		result = append(result, []int{nums[2], nums[3], nums[0], nums[1]})
		result = append(result, []int{nums[2], nums[3], nums[1], nums[0]})
		result = append(result, []int{nums[3], nums[0], nums[1], nums[2]})
		result = append(result, []int{nums[3], nums[0], nums[2], nums[1]})
		result = append(result, []int{nums[3], nums[1], nums[0], nums[2]})
		result = append(result, []int{nums[3], nums[1], nums[2], nums[0]})
		result = append(result, []int{nums[3], nums[2], nums[0], nums[1]})
		result = append(result, []int{nums[3], nums[2], nums[1], nums[0]})
	}
	return result
}
