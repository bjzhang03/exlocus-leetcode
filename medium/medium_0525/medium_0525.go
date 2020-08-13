package medium_0525

func findMaxLength(nums []int) int {
	result := 0
	if len(nums) > 0 {
		tmpsum := 0
		save := make([]int, len(nums))
		used := map[int][]int{}
		for i := 0; i < len(nums); i++ {
			// 将数组中的0换成1进行计算和
			if nums[i] == 0 {
				tmpsum = tmpsum - 1
			} else if nums[i] == 1 {
				tmpsum = tmpsum + 1
			}
			save[i] = tmpsum
			if _, ok := used[tmpsum]; ok {
				used[tmpsum] = append(used[tmpsum], i)
			} else if !ok {
				// 下标最小的在第一个
				used[tmpsum] = []int{i}
			}
			// 判断直接出现0的情况
			if tmpsum == 0 && result < i+1 {
				result = i + 1
			}
		}

		for i := len(save) - 1; i >= 0; i-- {
			if _, ok := used[save[i]]; ok && result < (i-used[save[i]][0]) {
				result = (i - used[save[i]][0])
			}
		}
	}
	return result
}
