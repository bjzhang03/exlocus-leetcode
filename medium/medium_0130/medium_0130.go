package medium_0130

type point struct {
	x int
	y int
}

func deepFirstSearch(board [][]byte, px int, py int, used map[point]bool, opoint *map[point]bool) {
	// 判断坐标没有超出
	if px >= 0 && py >= 0 && px < len(board) && py < len(board[0]) {
		// 当前的节点以前没有使用过
		ok := used[point{px, py}];
		ok2 := (*opoint)[point{px, py}];
		if !ok && !ok2 && board[px][py] == byte('O') {
			(*opoint)[point{px, py}] = true
			used[point{px, py}] = true

			deepFirstSearch(board, px+1, py, used, opoint)
			deepFirstSearch(board, px-1, py, used, opoint)
			deepFirstSearch(board, px, py-1, used, opoint)
			deepFirstSearch(board, px, py+1, used, opoint)

			used[point{px, py}] = false
		}
	}
}

func solve(board [][]byte) {
	// 直接从周围开始寻找即可
	if len(board) > 0 {
		opint := map[point]bool{}
		//寻找第一行
		for i := 0; i < len(board); i++ {
			deepFirstSearch(board, i, 0, map[point]bool{}, &opint)
			deepFirstSearch(board, i, len(board[0])-1, map[point]bool{}, &opint)
		}
		// 寻找第一列
		for i := 0; i < len(board[0]); i++ {
			deepFirstSearch(board, 0, i, map[point]bool{}, &opint)
			deepFirstSearch(board, len(board)-1, i, map[point]bool{}, &opint)
		}

		// 将所有的数据全部设置为'X'
		for i := 0; i < len(board); i++ {
			for j := 0; j < len(board[0]); j++ {
				board[i][j] = 'X'
			}
		}

		// 直接填充数据
		for key, _ := range opint {
			board[key.x][key.y] = 'O'
		}

	}
}
