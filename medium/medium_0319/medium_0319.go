package medium_0319

// https://www.cnblogs.com/grandyang/p/5100098.html
func bulbSwitch(n int) int {
	if n > 0 {
		result := 0
		istart := 1
		for istart*istart <= n {
			istart++
			result++
		}
		return result
	}
	return 0
}
