package medium_0390

// https://blog.csdn.net/qq_41410799/article/details/85284273
// https://www.cnblogs.com/wemo/p/10496189.html
// 这个题目是一个数学题目
// 我对这个题目的理解也不是很透彻
func lastRemaining(n int) int {
	result := 0

	if n == 1 {
		result = 1
	} else if n > 1 {
		result = 2 * (n/2 + 1 - lastRemaining(n/2))
	}
	return result
}
