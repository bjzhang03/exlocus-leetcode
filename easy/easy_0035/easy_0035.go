package easy_0035

func searchInsert(nums []int, target int) int {
	result := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] >= target {
			result = i
			break
		}
	}
	if result < 0 {
		result = len(nums)
	}
	return result
}
