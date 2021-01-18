package medium_0787

import (
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {
	return solve(n, flights, src, dst, K)
}

func solve(n int, flights [][]int, src int, dst int, K int) int {
	if n > 0 {
		// 初始化数据
		save := make([][]int, n)
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, n)
		}
		// 将所有的数据都初始化成math.max
		for i := 0; i < len(save); i++ {
			for j := 0; j < len(save[i]); j++ {
				save[i][j] = math.MaxInt32
			}
		}
		//将所有的节点数据进行更新
		for i := 0; i < len(flights); i++ {
			save[flights[i][0]][flights[i][1]] = flights[i][2]
		}
		// 使用广度优先遍历的方法来对数据进行更新
		current := []int{src}
		for len(current) > 0 && K >= 0 {
			// 每一层的时候使用,可以回去再找
			used := map[int]bool{}
			// 找出所有当前节点可以直接链接到的点，更新数据
			ncurrent := []int{}
			// 生成新的nsave
			nsave := make([][]int, n)
			for i := 0; i < len(nsave); i++ {
				nsave[i] = make([]int, n)
			}
			for i := 0; i < len(save); i++ {
				for j := 0; j < len(save[i]); j++ {
					nsave[i][j] = save[i][j]
				}
			}
			for i := 0; i < len(current); i++ {
				for j := 0; j < len(save[current[i]]); j++ {
					if nsave[src][j] > save[src][current[i]]+save[current[i]][j] {
						nsave[src][j] = save[src][current[i]] + save[current[i]][j]
					}
					// 当前的这些节点是可以达到的,并且是没有被使用过的
					if _, ok := used[j]; !ok {
						ncurrent = append(ncurrent, j)
						used[j] = true
					}
				}
			}
			current = append(current[:0], ncurrent...)
			// 更新save
			for i := 0; i < len(save); i++ {
				for j := 0; j < len(save); j++ {
					save[i][j] = nsave[i][j]
				}
			}
			K--
		}
		if save[src][dst] == math.MaxInt32 {
			return -1
		}
		return save[src][dst]
	}
	return 0
}
