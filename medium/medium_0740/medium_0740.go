package medium_0740

func deleteAndEarn(nums []int) int {
	if len(nums) > 0 {
		save := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			if _, ok := save[nums[i]]; ok {
				save[nums[i]]++
			} else if !ok {
				save[nums[i]] = 1
			}
		}

		return solve(save, &map[string]int{})
	}
	return 0
}

func solve(mnums map[int]int, save *map[string]int) int {
	result := 0
	if len(mnums) > 0 {
		for key, val := range mnums {
			nmnums := make(map[int]int)
			for nu, _ := range mnums {
				if nu != key {
					nmnums[nu] = mnums[key]
				}
			}
			if tres := solve(nmnums, save); result < key*val+tres {
				result = key*val + tres
			}
		}
	}
	return result
}
