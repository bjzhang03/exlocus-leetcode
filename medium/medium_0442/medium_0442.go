package medium_0442

// 这道题目应该还有其他的更优秀的解法
func findDuplicates(nums []int) []int {
	result := []int{}
	if len(nums) > 0 {
		save := make(map[int]int, len(nums))
		for i := 0; i < len(nums); i++ {
			if _, ok := save[nums[i]]; !ok {
				save[nums[i]] = 1
			} else if ok {
				save[nums[i]]++
			}
		}
		for key, val := range save {

			if val > 1 {
				result = append(result, key)
			}
		}
	}
	return result
}
