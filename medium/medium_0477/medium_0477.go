package medium_0477

// https://www.cnblogs.com/liujinhong/p/6206792.html
// 这个题目的技巧性很强
func totalHammingDistance(nums []int) int {
	result := 0
	for i := 0; i < 32; i++ {
		bitcount := 0
		for j := 0; j < len(nums); j++ {
			bitcount = bitcount + (nums[j]>>uint8(i))&1
		}
		result = result + bitcount*(len(nums)-bitcount)
	}
	return result
}
