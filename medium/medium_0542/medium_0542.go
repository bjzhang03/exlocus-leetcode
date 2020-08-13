package medium_0542

import (
	"strconv"
)

func updateMatrix(matrix [][]int) [][]int {
	if len(matrix) > 0 {
		// 先生成数组
		save := make([][]int, len(matrix))
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(matrix[0]))
		}
		// 将数组初始化
		used := map[string]bool{}
		queue := []*point{}
		for i := 0; i < len(save); i++ {
			for j := 0; j < len(save[i]); j++ {
				if matrix[i][j] == 0 {
					used[(&point{i, j}).toString()] = true
					queue = append(queue, &point{i, j})
				} else {
					// 题目已经说明数字个数不会大于10000
					save[i][j] = 100000
				}
			}
		}

		// bfs操作
		for len(queue) > 0 {
			curpoint := queue[0]
			// 上
			if checkpass(&matrix, used, &point{curpoint.x - 1, curpoint.y}) {
				save[curpoint.x-1][curpoint.y] = save[curpoint.x][curpoint.y] + 1
				queue = append(queue, &point{curpoint.x - 1, curpoint.y})
				used[(&point{curpoint.x - 1, curpoint.y}).toString()] = true
			}
			// 下
			if checkpass(&matrix, used, &point{curpoint.x + 1, curpoint.y}) {
				save[curpoint.x+1][curpoint.y] = save[curpoint.x][curpoint.y] + 1
				queue = append(queue, &point{curpoint.x + 1, curpoint.y})
				used[(&point{curpoint.x + 1, curpoint.y}).toString()] = true
			}
			// 左
			if checkpass(&matrix, used, &point{curpoint.x, curpoint.y - 1}) {
				save[curpoint.x][curpoint.y-1] = save[curpoint.x][curpoint.y] + 1
				queue = append(queue, &point{curpoint.x, curpoint.y - 1})
				used[(&point{curpoint.x, curpoint.y - 1}).toString()] = true
			}
			// 右
			if checkpass(&matrix, used, &point{curpoint.x, curpoint.y + 1}) {
				save[curpoint.x][curpoint.y+1] = save[curpoint.x][curpoint.y] + 1
				queue = append(queue, &point{curpoint.x, curpoint.y + 1})
				used[(&point{curpoint.x, curpoint.y + 1}).toString()] = true
			}
			queue = queue[1:]
		}
		return save
	}
	return nil
}

func checkpass(m *[][]int, used map[string]bool, p *point) bool {
	// 检查下标是否合法
	if !(p.x >= 0 && p.x < len(*m) && p.y >= 0 && p.y < len((*m)[0])) {
		return false
	}
	// 数据是已经使用过的数据
	if _, ok := used[p.toString()]; ok {
		return false
	}
	return true
}

type point struct {
	x int
	y int
}

// 题目里面已经说明矩阵不会大于10000，所以可以通过这种方式转化
func (p point) toString() string {
	return strconv.Itoa(p.x*10000 + p.y)
}
