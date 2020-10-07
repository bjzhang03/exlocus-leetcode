package easy_0231

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	} else {
		for n != 0 {
			if n&(0x1) == 1 && n > 1 {
				return false
			}
			n = n >> 1
		}
	}
	return true
}
