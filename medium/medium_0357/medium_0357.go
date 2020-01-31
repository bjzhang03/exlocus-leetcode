package medium_0357

func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	} else if n == 1 {
		return 9
	} else {
		return countNumbersWithUniqueDigits(n - 1)
	}

	return 0
}
