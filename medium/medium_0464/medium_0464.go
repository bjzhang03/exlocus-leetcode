package medium_0464

func canIWin(maxChoosableInteger int, desiredTotal int) bool {
	if maxChoosableInteger > 0 && desiredTotal > 0 && maxChoosableInteger*(maxChoosableInteger-1)/2 >= desiredTotal {
		choose := make(map[int]bool, maxChoosableInteger)
		for i := 1; i <= maxChoosableInteger; i++ {
			choose[i] = true
		}

		return winSolve(choose, desiredTotal, &map[string]bool{})
	}
	return false
}

func winSolve(choose map[int]bool, desired int, save *map[string]bool) bool {
	//fmt.Println(choose, desired)
	//fmt.Println(len(choose), choose, desired)
	// 对每一个可以使用的数据进行查找
	for key, val := range choose {
		if val {
			result := true
			// 如果还可以选择的话,则继续选,如果已经必胜了,则返回
			if desired-key > 0 {
				// 删掉数据
				delete(choose, key)
				for keys, vals := range choose {
					if vals {
						// 如果第二个人可以取走所有的数字,则直接返回false
						if keys >= desired-key {
							result = false
							break
						}
						delete(choose, keys)
						if !winSolve(choose, desired-key-keys, save) {
							result = false
						}
						// 回退现场
						choose[keys] = true
					}
					// 循环用到第二个人可以找的
				}
				// 回退现场
				choose[key] = true
			}
			// 选择当前的数据是否是必胜
			if result {
				// 选择当前数据则必胜
				return true
			}
		}
	}
	return false
}
