package medium_0371

// https://blog.csdn.net/Zhang_Yixuan_ss/article/details/80404812
func getSum(a int, b int) int {
	for a != 0 {
		tmp := (a & b) << 1
		b = a ^ b
		a = tmp
	}
	return b
}
