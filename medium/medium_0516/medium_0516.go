package medium_0516

func longestPalindromeSubseq(s string) int {
	if len(s) > 0 {
		save := make([][]int, len(s))
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(s))
		}
		// 对长度进行dp处理
		for lp := 0; lp <= len(s); lp++ {
			// 对坐标开始的长度进行dp操作
			for i := 0; i+lp < len(s); i++ {
				tmp := 0
				// 出现可能的回文
				if s[i] == s[i+lp] {
					if lp+i == i {
						tmp = 1
					} else {
						if i+1 <= lp+i-1 {
							tmp = save[i+1][i+lp-1] + 2
						} else {
							tmp = 2
						}
					}
				} else {
					// 在这个长度范围内找到最长的,此时就是当前的最大的值
					// 从前向后找
					for j := i; j <= i+lp; j++ {
						if tmp < save[j][i+lp] {
							tmp = save[j][i+lp]
						}
					}
					// 从后向前找
					for j := i + lp; j > i; j-- {
						if tmp < save[i][j] {
							tmp = save[i][j]
						}
					}
				}
				save[i][i+lp] = tmp
			}
		}
		return save[0][len(s)-1]
	}
	return 0
}
