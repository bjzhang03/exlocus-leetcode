package medium_0659

func isPossible(nums []int) bool {
	if len(nums) > 0 {
		save := [][]int{}
		for i := 0; i < len(nums); i++ {
			if len(save) == 0 {
				save = append(save, []int{nums[i]})
			} else {
				index := -1
				tmp := nums
				for j := 0; j < len(save); j++ {
					if save[j][len(save[j])-1]+1 == nums[i] && len(save[j]) < len(tmp) {
						tmp = save[j]
						index = j
					}
				}
				if len(tmp) == len(nums) {
					save = append(save, []int{nums[i]})
				} else {
					save[index] = append(save[index], nums[i])
				}
			}
		}

		result := true
		for i := 0; i < len(save); i++ {
			if len(save[i]) < 3 {
				result = false
			}
		}
		return result
	}
	return false
}
