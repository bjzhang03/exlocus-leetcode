package medium_0047

func hasslice(save [][]int, item []int) bool {
	result := false
	for i := 0; i < len(save); i++ {
		if len(save[i]) == len(item) {
			flag := true
			for j := 0; j < len(item); j++ {
				if save[i][j] != item[j] {
					flag = false
					break
				}
			}
			if flag {
				result = true
				break
			}
		}
	}
	return result
}

func dfs(nums []int, result *[][]int, save []int) {
	if len(nums) == 0 {
		if !hasslice(*result, save) {
			tmp := []int{}
			tmp = append(tmp, save...)
			*result = append(*result, tmp)
		}
		return
	} else if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			save = append(save, nums[i])
			nextnums := []int{}
			for j := 0; j < len(nums); j++ {
				if j != i {
					nextnums = append(nextnums, nums[j])
				}
			}
			dfs(nextnums, result, save)
			save = save[:len(save)-1]
		}
	}
}
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		save := []int{}
		dfs(nums, &result, save)
	}
	return result
}
