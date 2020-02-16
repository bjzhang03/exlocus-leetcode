package medium_0518

// https://www.jianshu.com/p/4557890d1bbc
// 这个题目我出现了建模错误,我还是对dp不是特别熟悉啊
func change(amount int, coins []int) int {
	if amount == 0 {
		return 1
	} else if amount >= 0 && len(coins) > 0 {
		save := make([][]int, len(coins))
		for i := 0; i < len(coins); i++ {
			save[i] = make([]int, amount+1)
		}
		for i := 0; i < len(coins); i++ {
			// 只是用到i之前的硬币共有多少中兑换方法
			for j := 0; j <= amount; j++ {
				tmp := 0
				for k := 0; j-k >= 0; k = k + coins[i] {
					if j == k {
						tmp++
					} else {
						// 使用前面的货币可以兑换的数据
						if i-1 >= 0 {
							tmp = tmp + save[i-1][j-k]
						}
					}
				}
				save[i][j] = tmp
			}
		}
		return save[len(coins)-1][amount]
	}
	return 0
}
