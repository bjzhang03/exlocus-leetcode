package medium_0079

type poisition struct {
	x int
	y int
}

func deepFirstSearch(board [][]byte, word string, wi int, used map[poisition]bool, px int, py int) bool {
	if wi >= len(word) {
		return true
	} else if uint8(board[px][py]) == word[wi] {
		// 把当前的数据标记为已经使用过的数据
		used[poisition{px, py}] = true
		hasnext := false
		// 没有使用过的才会继续寻找
		if ok := used[poisition{px - 1, py}]; px-1 >= 0 && !ok {
			hasnext = true
			nextFind := deepFirstSearch(board, word, wi+1, used, px-1, py)
			if nextFind {
				return true
			}
		}

		if ok := used[poisition{px, py - 1}]; py-1 >= 0 && !ok {
			hasnext = true
			nextFind := deepFirstSearch(board, word, wi+1, used, px, py-1)
			if nextFind {
				return true
			}
		}

		if ok := used[poisition{px + 1, py}]; px+1 < len(board) && !ok {
			hasnext = true
			nextFind := deepFirstSearch(board, word, wi+1, used, px+1, py)
			if nextFind {
				return true
			}
		}

		if ok := used[poisition{px, py + 1}]; py+1 < len(board[0]) && !ok {
			hasnext = true
			nextFind := deepFirstSearch(board, word, wi+1, used, px, py+1)
			if nextFind {
				return true
			}
		}
		
		// 修复状态数据
		used[poisition{px, py}] = false
		// 如果这个已经没有下一个节点了,则需要进行特殊的判断
		if !hasnext && wi+1 >= len(word) {
			return true
		}

	}
	return false
}

func exist(board [][]byte, word string) bool {
	// 找出所有的可能的开始的点
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if deepFirstSearch(board, word, 0, map[poisition]bool{}, i, j) {
				return true
			}

		}
	}
	return false
}
