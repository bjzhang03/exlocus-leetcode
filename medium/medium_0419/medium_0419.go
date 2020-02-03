package medium_0419

import (
	"reflect"
)

func countBattleships(board [][]byte) int {
	result := 0
	if len(board) > 0 {
		// 默认所有的数据都是0
		save := make([][]byte, len(board))
		for i := 0; i < len(board); i++ {
			save[i] = make([]byte, len(board[0]))
		}
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[i]); j++ {
				if save[i][j] == 0 && board[i][j] == 'X' {
					used := [][2]int{}
					xCluster(board, i, j, &used)
					if isBattleship(used) {
						result++
					}
					// 污染掉所有相关的数据,避免重复寻找
					for i := 0; i < len(used); i++ {
						save[used[i][0]][used[i][1]] = 1
					}
				}
			}
		}
	}
	return result
}

// 找到所有的cluster
func xCluster(board [][]byte, sti, eni int, used *[][2]int) {
	if checkCondition(board, sti, eni, *used) && board[sti][eni] == 'X' {
		(*used) = append(*used, [2]int{sti, eni})
		xCluster(board, sti-1, eni, used)
		xCluster(board, sti+1, eni, used)
		xCluster(board, sti, eni-1, used)
		xCluster(board, sti, eni+1, used)
	}
}

// 前置检查条件
func checkCondition(board [][]byte, ix, iy int, used [][2]int) bool {
	result := true
	if !(ix >= 0 && iy >= 0 && ix < len(board) && iy < len(board[0])) {
		result = false
	} else {
		for i := 0; i < len(used); i++ {
			if reflect.DeepEqual([2]int{ix, iy}, used[i]) {
				result = false
				break
			}
		}
	}
	return result
}

func isBattleship(ship [][2]int) bool {
	result := true
	if len(ship) > 1 {
		// 判断是不是所有的行都是一样的
		result = false
		if func(slices [][2]int) bool {
			// 所有的行都是一样的
			result := true
			for i := 1; i < len(slices); i++ {
				if slices[i][0] != slices[i-1][0] {
					result = false

				}
			}
			return result
		}(ship) ||
			func(slices [][2]int) bool {
				// 所有的列都是一样的
				result := true
				for i := 1; i < len(slices); i++ {
					if slices[i][1] != slices[i-1][1] {
						result = false

					}
				}
				return result
			}(ship) {
			result = true
		}
	}
	return result
}
