package medium_0503

// 直接暴力破解,竟然也可以
func nextGreaterElements(nums []int) []int {
	if len(nums) > 0 {
		result := []int{}
		save := append([]int{}, nums...)
		save = append(save, nums...)
		for i := 0; i < len(nums); i++ {
			tmp := nums[i]
			// 从当前数字开始向后寻找,找到第一个比自己大的数据即可
			for j := i; j < i+len(nums); j++ {
				if tmp < save[j] {
					tmp = save[j]
					break
				}
			}
			// 判断是不是自己,是自己表明没有找到更大的数字
			if tmp == nums[i] {
				result = append(result, -1)
			} else {
				result = append(result, tmp)
			}
		}
		return result
	}
	return nil
}
