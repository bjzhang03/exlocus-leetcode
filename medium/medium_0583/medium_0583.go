package medium_0583

//这里的接替思路是按照lcs来进行动态规划求解的
func minDistance(word1 string, word2 string) int {
	if len(word1) > 0 && len(word2) > 0 {
		// 初始化的时候,所有的数据都是0
		save := make([][]int, len(word1)+1)
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(word2)+1)
		}
		// 动态规划转移方程
		for i := 1; i <= len(word1); i++ {
			for j := 1; j <= len(word2); j++ {
				if word1[i-1] == word2[j-1] {
					save[i][j] = save[i-1][j-1] + 1
				} else if word1[i-1] != word2[j-1] {
					if save[i-1][j] > save[i][j-1] {
						save[i][j] = save[i-1][j]
					} else {
						save[i][j] = save[i][j-1]
					}
				}
			}
		}
		//返回一共需要删除的字符的个数
		return len(word1) + len(word2) - 2*save[len(word1)][len(word2)]
	} else if len(word1) == 0 {
		return len(word2)
	} else if len(word2) == 0 {
		return len(word1)
	}
	return 0
}
