package medium_0150

import "strconv"

func evalRPN(tokens []string) int {
	nums := []int{}
	for _, val := range tokens {
		if val == "*" {
			left := nums[len(nums)-2]
			right := nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, left*right)
		} else if val == "/" {
			// 不用担心除数为0的问题
			left := nums[len(nums)-2]
			right := nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, left/right)
		} else if val == "+" {
			left := nums[len(nums)-2]
			right := nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, left+right)
		} else if val == "-" {
			left := nums[len(nums)-2]
			right := nums[len(nums)-1]
			nums = nums[:len(nums)-2]
			nums = append(nums, left-right)
		} else {
			num, _ := strconv.Atoi(val)
			nums = append(nums, num)
		}
	}
	return nums[0]
}
