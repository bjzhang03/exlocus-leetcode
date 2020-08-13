package medium_0523

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) >= 2 {
		// 计算sum数据,并存储下来
		save := make([]int, len(nums)+1)
		numsum := 0
		for i := 0; i < len(nums); i++ {
			numsum = numsum + nums[i]
			save[i+1] = numsum
		}

		// 对每一个数据进行判断
		for i := 0; i < len(save); i++ {
			for j := i + 2; j < len(save); j++ {
				// 对特殊的k的数据进行分类处理
				if k != 0 && (save[j]-save[i])%k == 0 {
					return true
				} else if k == 0 && (save[j]-save[i]) == 0 {
					return true
				}

			}
		}
	}
	return false
}
