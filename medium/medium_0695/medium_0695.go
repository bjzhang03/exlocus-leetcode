package medium_0695

func maxAreaOfIsland(grid [][]int) int {
	if len(grid) > 0 {
		// used 被所有的循环共享,来标记是否被使用过了
		used := make([][]int, len(grid))
		for i := 0; i < len(used); i++ {
			used[i] = make([]int, len(grid[i]))
		}
		// 挨个数据进行深度优先遍历查找
		result := 0
		for i := 0; i < len(grid); i++ {
			for j := 0; j < len(grid[i]); j++ {
				if used[i][j] == 0 && grid[i][j] == 1 {
					if current := deepFirstSearct(&grid, &used, []int{i, j}); result < current {
						result = current
					}
				}
			}
		}
		return result
	}
	return 0
}

// 深度优先搜索
func deepFirstSearct(grid *[][]int, used *[][]int, point []int) int {
	if func(g *[][]int, u *[][]int, p []int) bool {
		// 判断数据下标是否合格
		if !(p[0] >= 0 && p[0] < len(*g) && p[1] >= 0 && p[1] < len((*g)[0])) {
			return false
		}
		// 判断数据是否以前被标记过,标记过的话,则不再进行处理
		if ((*u)[p[0]][p[1]]) > 0 {
			return false
		}
		return true
	}(grid, used, point) {
		result := 0
		if (*grid)[point[0]][point[1]] == 1 {
			result = result + 1
			// 这里最好新建used给下一次的循环使用
			(*used)[point[0]][point[1]] = 1
			result += deepFirstSearct(grid, used, []int{point[0] - 1, point[1]})
			result += deepFirstSearct(grid, used, []int{point[0] + 1, point[1]})
			result += deepFirstSearct(grid, used, []int{point[0], point[1] - 1})
			result += deepFirstSearct(grid, used, []int{point[0], point[1] + 1})
		}
		return result
	}
	return 0
}
