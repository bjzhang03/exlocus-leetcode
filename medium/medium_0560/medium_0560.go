package medium_0560

func subarraySum(nums []int, k int) int {
	if len(nums) > 0 {
		count := 0
		for i := 0; i < len(nums); i++ {
			tmps := 0
			for j := i; j < len(nums); j++ {
				tmps = tmps + nums[j]
				if tmps == k {
					count++
				}
			}
		}
		return count
	}
	return 0
}
