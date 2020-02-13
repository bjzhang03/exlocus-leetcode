package medium_0494

import "strconv"

func findTargetSumWays(nums []int, S int) int {
	return solve(nums, len(nums)-1, S, &map[string]int{})
}

func solve(nums []int, sti int, S int, used *map[string]int) int {
	result := 0
	if _, ok := (*used)[toString(sti, S)]; ok {
		return (*used)[toString(sti, S)]
	} else if func(nu []int, si int, target int) bool {
		// 先进行前提条件判断
		// 先判断是否还有数据可以选择
		if len(nu) <= 0 {
			return false
		}
		// 先判断下标的范围
		if si >= len(nu) || si < 0 {
			return false
		}
		return true
	}(nums, sti, S) {
		// 处理特殊的例子,只有一个数据的时候
		if sti == 0 {
			if nums[0] == S {
				result++
			}
			if nums[0] == -S {
				result++
			}
		} else {
			result = result + solve(nums, sti-1, S-nums[sti], used) + solve(nums, sti-1, S+nums[sti], used)
		}
		(*used)[toString(sti, S)] = result
	}
	return result
}

func toString(sti, S int) string {
	return "(" + strconv.Itoa(sti) + "," + strconv.Itoa(S) + ")"
}
