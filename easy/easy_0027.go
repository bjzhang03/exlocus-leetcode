package easy

func removeElement(nums []int, val int) int {
	i := 0
	j := 0
	for ; i < len(nums) && j < len(nums); i++ {
		if nums[j] == val && nums[i] != val {
			nums[j] = nums[i]
			nums[i] = val
			j++
		} else if nums[j] != val {
			j++
		}
	}
	return j
}
