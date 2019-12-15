package medium_0289

// 判断index是否合法
func legalindex(board *[][]int, x, y int) bool {
	if x >= 0 && x < len(*board) && y >= 0 && y < len((*board)[0]) {
		return true
	}
	return false
}

// 判断一个point周围有多少死细胞活细胞
func arround(board [][]int, x, y int) (int, int) {
	livecount := 0
	deadcount := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			// 不等于自身 + 合法的index
			if legalindex(&board, i, j) {
				if i == x && j == y {
					// 啥都不做
				} else if board[i][j] == 0 {
					deadcount++
				} else if board[i][j] == 1 {
					livecount++
				}
			}
		}
	}
	return livecount, deadcount
}

type point struct {
	x int
	y int
}

func gameOfLife(board [][]int) {
	if len(board) > 0 {
		livepoint := map[point]bool{}
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				// 判断当前的点周围有多少d和l
				// d周围有三个l
				if board[i][j] == 0 {
					live, _ := arround(board, i, j)
					if live == 3 {
						livepoint[point{i, j}] = true
					}
				} else if board[i][j] == 1 {
					live, dead := arround(board, i, j)
					if live < 2 {
						// live小于2将会死去
					} else if live == 2 || live == 3 {
						// live为2或者3将会活着
						livepoint[point{i, j}] = true
					} else if dead > 3 {
						// 如果dead大于3将会dead
					}
				}
			}
		}

		// 将board中所有的数据全部设置为0
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				board[i][j] = 0
			}
		}
		// 填充活细胞
		for key := range livepoint {
			board[key.x][ key.y] = 1
		}
	}
}
