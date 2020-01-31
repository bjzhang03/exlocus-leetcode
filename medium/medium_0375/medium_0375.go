package medium_0375

import (
	"math"
)

// https://blog.csdn.net/abc15766228491/article/details/82931660
// https://blog.csdn.net/qq508618087/article/details/51991143
// 这个题目不是很好理解,需要细细琢磨
// 仔细阅读题目
func getMoneyAmount(n int) int {
	result := 0
	// 构建数组
	save := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		save[i] = make([]int, n+1)
	}

	// 初始化数组数据
	for i := 0; i <= n; i++ {
		// 自己到自己的数据为0
		save[i][i] = 0
		// 只有两个的时候，自己到下一个数据的话，选择自己
		if i+1 <= n {
			save[i][i+1] = i
		}
	}

	// 按照len长度进行dp操作
	for len := 2; len <= n; len++ {
		for i := 1; i <= n; i++ {
			if i+len <= n {
				tmp := math.MaxInt64
				for j := i + 1; j <= i+len && j <= n && i+len <= n; j++ {
					// 需要确认当前找的不是最后面的一个
					if j+1 <= n {
						// 左边寻找付出的代价更高
						if save[j+1][i+len] > save[i][j-1] {
							if tmp > save[j+1][i+len]+j {
								tmp = save[j+1][i+len] + j
							}
						} else {
							// 右边寻找付出的代价更高
							if tmp > save[i][j-1]+j {
								tmp = save[i][j-1] + j
							}
						}
					}
					// 找出两边寻找的最大代价的最小值
				}
				save[i][i+len] = tmp
			}
		}
	}
	result = save[1][n]

	return result
}
