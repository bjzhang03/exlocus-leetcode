package medium_0712

// https://blog.csdn.net/qq_41822647/article/details/86660550
// 应该是我自己做错了,后来我参考了上面的方案
// 这个题目时最长公共子串的一个变种的题目,我之前使用先找出最长公共子串,
// 然后再进行判断的操作,最终没有通过,所以比较坑，我参考了上面的方案
// 我觉得还是我自己没有冷静思考和对最长公共子串的理解力不足
func minimumDeleteSum(s1, s2 string) int {
	if len(s1) > 0 && len(s2) > 0 {
		save := make([][]int, len(s1)+1)
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(s2)+1)
		}

		for i := 0; i < len(s1); i++ {
			save[i+1][0] = save[i][0] + int(s1[i])
		}
		for j := 0; j < len(s2); j++ {
			save[0][j+1] = save[0][j] + int(s2[j])
		}

		for i := 0; i < len(s1); i++ {
			for j := 0; j < len(s2); j++ {
				if s1[i] == s2[j] {
					save[i+1][j+1] = save[i][j]
				} else {
					if (save[i+1][j] + int(s2[j])) > (save[i][j+1] + int(s1[i])) {
						save[i+1][j+1] = save[i][j+1] + int(s1[i])
					} else {
						save[i+1][j+1] = save[i+1][j] + int(s2[j])
					}
				}
			}
		}
		return save[len(s1)][len(s2)]
	}
	return 0
}
