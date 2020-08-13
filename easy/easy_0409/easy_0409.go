package easy_0409

func longestPalindrome(s string) int {
	if len(s) > 0 {
		// 统计个数出现的次数的数据
		save := map[uint8]int{}
		for i := 0; i < len(s); i++ {
			if _, ok := save[s[i]]; ok {
				save[s[i]] = save[s[i]] + 1
			} else if !ok {
				save[s[i]] = 1
			}
		}
		// 构造结果
		maxod := 0
		result := 0
		for _, val := range save {
			// 偶数个直接拿进来
			if val%2 == 0 {
				result = result + val
			} else if maxod < val {
				// 奇数个,则减一个再拿进来
				maxod = 1
				result = result + val - 1
			}
		}
		return result + maxod
	}
	return 0
}
