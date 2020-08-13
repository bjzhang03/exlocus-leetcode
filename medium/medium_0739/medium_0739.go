package medium_0739

func dailyTemperatures(T []int) []int {
	if len(T) > 0 {
		// 新建一个map存储所有的下标的数据
		msave := make(map[int][]int)
		for i := 0; i < len(T); i++ {
			if _, ok := msave[T[i]]; ok {
				// 这里的下标是按照顺序添加进去的
				msave[T[i]] = append(msave[T[i]], i)
			} else if !ok {
				msave[T[i]] = []int{i}
			}
		}

		// 从map中找出符合条件的数字
		result := []int{}
		for i := 0; i < len(T); i++ {
			// 遍历所有的比当前的数字大的数据
			// 30001是题目中告知的最大的T的长度,100是题目中告知的温度的范围
			count := 30001
			for val := T[i] + 1; val <= 100; val++ {
				// 挨个数字查找
				if _, ok := msave[val]; ok {
					for _, ival := range msave[val] {
						// 因为下标是按照升序添加进去的,所以第一个即可
						if ival > i && count > (ival-i) {
							count = ival - i
							break
						}
					}
				}
			}
			if count != 30001 {
				result = append(result, count)
			} else {
				result = append(result, 0)
			}
		}
		return result
	}
	return nil
}
