package medium_0554

func leastBricks(wall [][]int) int {
	if len(wall) > 0 {
		used := map[int]int{}
		// 计算砖墙的长度
		brickslen := 0
		for i := 0; i < len(wall[0]); i++ {
			brickslen = brickslen + wall[0][i]
		}

		brickscount := 0
		for i := 0; i < len(wall); i++ {
			tmp := 0
			for j := 0; j < len(wall[i]); j++ {
				tmp = tmp + wall[i][j]
				if _, ok := used[tmp]; ok {
					used[tmp]++
				} else if !ok {
					used[tmp] = 1
				}
				// 顺便判断当前的数据是否是最大值
				if tmp != brickslen && used[tmp] > brickscount {
					brickscount = used[tmp]
				}
			}
		}
		return len(wall) - brickscount
	}
	return 0
}
