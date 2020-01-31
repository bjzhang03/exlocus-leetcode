package medium_0377

func combinationSum4(nums []int, target int) int {
	return solve(nums, target, map[int]int{})
}

//因为是在一个package里面,所以需要新建一个函数来隔离全局变量的数据
func solve(nums []int, target int, save map[int]int) int {
	result := 0
	if key, ok := save[target]; ok {
		result = key
	} else if len(nums) > 0 && target > 0 {
		for i := 0; i < len(nums); i++ {
			result = result + solve(nums, target-nums[i], save)
			if nums[i] == target {
				result++
			}
		}
		save[target] = result
	}
	return result
}
