package medium_0376

func wiggleMaxLength(nums []int) int {
	result := 0
	save := make([][2]int, len(nums))
	for i := 0; i < len(nums); i++ {
		current := 0
		count := 1
		for j := i - 1; j >= 0; j-- {
			// 这里的等于是处理一种特殊情况 [3,3,3,2,5]
			if save[j][0] <= 0 {
				// 当前的数据需要比j的数据大
				if nums[i] > nums[j] && count < save[j][1]+1 {
					current = nums[i] - nums[j]
					count = save[j][1] + 1
				}
			}

			// 这里的等于是处理一种特殊情况 [3,3,3,2,5]
			if save[j][0] >= 0 {
				// 当前的数据需要比j的数据小
				if nums[i] < nums[j] && count < save[j][1]+1 {
					current = nums[i] - nums[j]
					count = save[j][1] + 1
				}
			}
		}
		save[i][0] = current
		save[i][1] = count
	}

	for i := 0; i < len(save); i++ {
		if result < save[i][1] {
			result = save[i][1]
		}
	}
	return result
}
