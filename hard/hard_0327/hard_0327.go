package hard_0327

func countRangeSum(nums []int, lower int, upper int) int {
	result := 0
	if len(nums) > 0 {
		// 计算临时处理的数据
		sums := make([]int, len(nums)+1)
		tmp := 0
		for i := 0; i < len(nums); i++ {
			tmp = tmp + nums[i]
			sums[i+1] = tmp
		}

		for i := len(sums) - 1; i > 0; i-- {
			for j := i - 1; j >= 0; j-- {
				if sums[i]-sums[j] >= lower && sums[i]-sums[j] <= upper {
					result++
				}
			}
		}
	}
	return result
}
