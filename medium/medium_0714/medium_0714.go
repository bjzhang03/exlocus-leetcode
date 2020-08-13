package medium_0714

//这个题目没有ac
func maxProfit(prices []int, fee int) int {
	if len(prices) > 0 {
		return solve(prices, fee)
	}
	return 0
}

func solve(prices []int, fee int) int {
	if len(prices) > 0 {
		save := make([][]int, len(prices))
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(prices))
		}
		// 使用动态规划的方式来计算当前数据的最优值
		for plen := 1; plen < len(prices); plen++ {
			for i := 0; i+plen < len(prices); i++ {
				stmp := 0
				if stmp < prices[i+plen]-prices[i]-fee {
					stmp = prices[i+plen] - prices[i] - fee
				}
				// 数据分割进行计算
				for k := i + 1; k < i+plen; k++ {
					if stmp < save[i][k]+save[k][i+plen] {
						stmp = save[i][k] + save[k][i+plen]
					}
				}
				save[i][i+plen] = stmp
			}
		}
		return save[0][len(prices)-1]
	}
	return 0
}
