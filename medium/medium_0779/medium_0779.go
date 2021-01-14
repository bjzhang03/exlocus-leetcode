package medium_0779

/*直接根据前面的一个来计算当前的数据的正确的结果
这个方法是自己想出来的,不是百度来的*/
func kthGrammar(N int, K int) int {
	if N == 1 && K == 1 {
		return 0
	} else if N > 1 && K > 1 {
		if pres := kthGrammar(N-1, (K+1)/2); pres == 0 {
			if K%2 == 0 {
				return 1
			} else {
				return 0
			}
		} else if pres == 1 {
			if K%2 == 0 {
				return 0
			} else {
				return 1
			}
		}
	}
	return 0
}
