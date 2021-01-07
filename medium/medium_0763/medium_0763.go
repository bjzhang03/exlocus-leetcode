package medium_0763

func partitionLabels(S string) []int {
	return solve(S)
}

func solve(str string) []int {
	right := map[uint8]int{}
	for i := 0; i < len(str); i++ {
		if _, ok := right[str[i]]; !ok {
			right[str[i]] = 1
		} else if ok {
			right[str[i]]++
		}
	}

	left := map[uint8]int{}
	for i := 0; i < len(str); i++ {
		// 更新右边的数据
		right[str[i]]--
		if right[str[i]] == 0 {
			delete(right, str[i])
		}
		// 更新左边的数据
		if _, ok := left[str[i]]; !ok {
			left[str[i]] = 1
		} else if ok {
			left[str[i]]++
		}
		// 判断两个map是否相交
		if !mapJoin(left, right) {
			return append([]int{i + 1}, solve(str[i+1:])...)
		}
	}
	return nil
}

func mapJoin(a, b map[uint8]int) bool {
	for key, _ := range a {
		if _, ok := b[key]; ok {
			return true
		}
	}
	for key, _ := range b {
		if _, ok := a[key]; ok {
			return true
		}
	}
	return false
}
