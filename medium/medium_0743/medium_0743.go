package medium_0743

// 这个其实就是最短路径算法
func networkDelayTime(times [][]int, N int, K int) int {
	if len(times) > 0 && N > 0 {
		// 先将两个图转化正一张表
		save := make([][]int, N+1)
		for i := 0; i <= N; i++ {
			save[i] = make([]int, N+1)
		}
		// 构造数据为最大值
		for i := 0; i <= N; i++ {
			for j := 0; j <= N; j++ {
				save[i][j] = 6000000
			}
		}
		// 将图转化成表
		for i := 0; i < len(times); i++ {
			save[times[i][0]][times[i][1]] = times[i][2]
			save[times[i][1]][times[i][0]] = times[i][2]
		}

		visit := []int{K}
		for len(visit) > 0 {
			for i := 0; i <= N; i++ {

			}
		}
	}
	return -1
}
