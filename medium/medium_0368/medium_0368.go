package medium_0368

import "sort"

func largestDivisibleSubset(nums []int) []int {
	result := []int{}
	save := make([][]int, len(nums))
	// 先对数据进行排序操作
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		tmp := []int{}
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(tmp) < len(save[j]) {
				tmp = tmp[0:0]
				tmp = append(tmp, save[j]...)
			}
		}
		tmp = append(tmp, nums[i])
		save[i] = append(save[i], tmp...)

		if len(result) < len(tmp) {
			result = result[0:0]
			result = append(result, tmp...)
		}
	}

	return result
}
