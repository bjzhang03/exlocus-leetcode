package medium_0565

func arrayNesting(nums []int) int {
	if len(nums) > 0 {
		used := map[int]bool{}
		result := 0
		for i := 0; i < len(nums); i++ {
			if _, ok := used[nums[i]]; !ok {
				if curlen := getCircle(i, &nums, &used); result < curlen {
					result = curlen
				}
			}
		}
		return result
	}
	return 0
}

func getCircle(start int, nums *[]int, used *map[int]bool) int {
	if start >= 0 && start < len(*nums) {
		current := map[int]bool{}
		for {
			if _, ok := current[(*nums)[start]]; ok {
				break
			} else {
				current[(*nums)[start]] = true
				start = (*nums)[start]
			}
		}

		// 更新used的数据
		for key, val := range current {
			(*used)[key] = val
		}
		return len(current)
	}
	return 0
}
