package medium_0343

func integerBreak(n int) int {
	result := 0
	if n > 0 {
		// 题目中已经假设n小于等于52了
		save := make([]int, 100)
		save[1] = 1
		save[2] = 1

		// 直接从3开始
		for i := 3; i <= n; i++ {
			current := 0
			for j := 1; j < i; j++ {
				if current < save[j]*save[i-j] {
					current = save[j] * save[i-j]
				}
				// 特殊情况计算,直接两部分组成
				if current < j*(i-j) {
					current = j * (i - j)
				}
				// 当前的结果不拆分
				if current < save[j]*(i-j) {
					current = save[j] * (i - j)
				}
			}
			save[i] = current
		}
		result = save[n]
	}
	return result
}
