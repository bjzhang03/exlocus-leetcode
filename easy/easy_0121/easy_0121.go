package easy_0121

func maxProfit(prices []int) int {
	result := 0
	if len(prices) > 0 {
		for i := 0; i < len(prices); i++ {
			for j := i + 1; j < len(prices); j++ {
				if result < prices[j]-prices[i] {
					result = prices[j] - prices[i]
				}
			}
		}
	}
	return result
}
