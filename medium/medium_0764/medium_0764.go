package medium_0764

func orderOfLargestPlusSign(N int, mines [][]int) int {
	return solve(N, mines)
}

func solve(N int, mines [][]int) int {
	result := 0
	// 构建表
	table := make([][]int, N)
	for i := 0; i < len(table); i++ {
		table[i] = make([]int, N)
	}
	// 将有标记的位置标记为1
	for i := 0; i < len(mines); i++ {
		table[mines[i][0]][mines[i][1]] = 1
	}
	// 再每一个位置检查是否有加号
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			if tres := plusSign(i, j, &table); result < tres {
				result = tres
			}
		}
	}
	return result
}

func plusSign(ix, iy int, table *[][]int) int {
	result := 0
	if (*table)[ix][iy] == 0 {
		for ix+result < len(*table) && ix-result >= 0 && iy+result < len(*table) && iy-result >= 0 {
			if (*table)[ix+result][iy] == 0 && (*table)[ix-result][iy] == 0 && (*table)[ix][iy+result] == 0 && (*table)[ix][iy-result] == 0 {
				result++
			} else {
				break
			}
		}
	}
	return result
}
