package medium_0421

// 这么差的算法竟然ac了
func findMaximumXOR(nums []int) int {
	result := 0
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			for j := i; j < len(nums); j++ {
				if result < nums[i]^nums[j] {
					result = nums[i] ^ nums[j]
				}
			}
		}
	}
	return result
}
