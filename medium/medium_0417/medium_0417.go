package medium_0417

import "fmt"

var xmul = 1000

func pacificAtlantic(matrix [][]int) [][]int {
	result := [][]int{}
	if len(matrix) > 0 {
		lpx := 0
		lpy := 0
		rbx := len(matrix) - 1
		rby := len(matrix[0]) - 1
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[0]); j++ {
				// 四个方向检查即可
				if passCheck(i-1, j, &matrix, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, lpx, lpy, matrix[i][j], i-1, j, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, rbx, rby, matrix[i][j], i-1, j, &map[int]bool{i*xmul + j: true}) {
					result = append(result, []int{i, j})

				} else if passCheck(i+1, j, &matrix, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, lpx, lpy, matrix[i][j], i+1, j, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, rbx, rby, matrix[i][j], i+1, j, &map[int]bool{i*xmul + j: true}) {
					result = append(result, []int{i, j})
				} else if passCheck(i, j-1, &matrix, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, lpx, lpy, matrix[i][j], i, j-1, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, rbx, rby, matrix[i][j], i, j-1, &map[int]bool{i*xmul + j: true}) {
					result = append(result, []int{i, j})
				} else if passCheck(i, j+1, &matrix, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, lpx, lpy, matrix[i][j], i, j+1, &map[int]bool{i*xmul + j: true}) &&
					deepFirstSearch(&matrix, rbx, rby, matrix[i][j], i, j+1, &map[int]bool{i*xmul + j: true}) {
					result = append(result, []int{i, j})
				}
			}
		}
	}
	fmt.Println(result)
	return result
}

func deepFirstSearch(matrix *[][]int, tarx, tary int, tarnum int, curx, cury int, save *map[int]bool) bool {
	fmt.Println("num = ", tarnum, tarx, tary, curx, cury, *save)
	if passCheck(curx, cury, matrix, save) {
		if curx == tarx && cury == tary {
			return true
		} else if (*matrix)[curx][cury] <= tarnum {
			// 上下左右各个方向查找
			(*save)[curx*xmul+cury] = true
			// 向上查找
			if passCheck(curx-1, cury, matrix, save) && deepFirstSearch(matrix, tarx, tary, tarnum, curx-1, cury, save) {
				return true
			}
			// 向下查找
			if passCheck(curx+1, cury, matrix, save) && deepFirstSearch(matrix, tarx, tary, tarnum, curx+1, cury, save) {
				return true
			}
			// 向左查找
			if passCheck(curx, cury-1, matrix, save) && deepFirstSearch(matrix, tarx, tary, tarnum, curx, cury-1, save) {
				return true
			}
			// 向右查找
			if passCheck(curx, cury+1, matrix, save) && deepFirstSearch(matrix, tarx, tary, tarnum, curx, cury+1, save) {
				return true
			}
			// 还原数据现场
			(*save)[curx*xmul+cury] = false
		}
	}
	return false
}

func passCheck(x, y int, matrix *[][]int, save *map[int]bool) bool {
	result := true
	if !(x >= 0 && x < len(*matrix) && y >= 0 && y < len((*matrix)[0])) {
		// 判断下标是否正确
		result = false
	} else if key, ok := (*save)[xmul*x+y]; ok && key {
		// 判断当前的数据是否是已经使用过了的数据
		result = false
	}

	fmt.Println("check: ", x, y, *save, result)

	return result
}
