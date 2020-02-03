package medium_0137

// https://blog.csdn.net/zsjwish/article/details/80038434
func singleNumber(nums []int) int {
	result := 0
	if len(nums) > 0 {
		// 默认的int长度是64
		for i := 0; i < 64; i++ {
			sum := 0
			for _, num := range nums {
				sum = sum + (num>>uint8(i))&1
			}
			if sum%3 != 0 {
				result = result | 1<<uint8(i)
			}
		}
	}
	return result
}
