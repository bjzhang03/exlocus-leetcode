package medium_0743

import (
	"math"
)

// 这个其实就是最短路径算法
func networkDelayTime(times [][]int, N int, K int) int {
	if len(times) > 0 && N > 0 {
		// 将图变换成表
		save := make([][]int, N+1)
		for i := 0; i <= N; i++ {
			save[i] = make([]int, N+1)
		}
		for i := 0; i <= N; i++ {
			for j := 0; j < len(save[i]); j++ {
				save[i][j] = math.MaxInt32
			}
		}
		for i := 0; i < len(times); i++ {
			save[times[i][0]][times[i][1]] = times[i][2]
		}
		save[K][K] = 0

		// 最短路径算法
		visit := []int{K}
		visited := map[int]bool{K: true}
		used := map[int]bool{}
		for len(visit) > 0 {
			// 找到最短的数据
			plen := math.MaxInt32
			index := 0
			for i := 0; i < len(visit); i++ {
				if plen > save[K][visit[i]] {
					plen = save[K][visit[i]]
					index = i
				}
			}
			// 将最短的数据标记为已经使用过的数据
			used[visit[index]] = true
			current := visit[index]
			// 更细visit的数据
			visit = append(visit[:index], visit[index+1:]...)

			// 更新一遍所有的数据
			for i := 0; i <= N; i++ {
				// 当前的数据是可以达到的
				if save[K][current]+save[current][i] < math.MaxInt32 {
					if _, ok := visited[i]; !ok {
						visit = append(visit, i)
						visited[i] = true
					}
				}
				// 当前的数据需要更新
				if save[K][i] > save[K][current]+save[current][i] {
					save[K][i] = save[K][current] + save[current][i]
				}
			}
		}
		// 寻找随后的数据
		if len(used) < N {
			return -1
		} else if len(used) == N {
			// 获取最终的数据
			plen := math.MinInt32
			for i := 0; i <= N; i++ {
				if save[K][i] < math.MaxInt32 && plen < save[K][i] {
					plen = save[K][i]
				}
			}
			return plen
		}
	}
	return -1
}
