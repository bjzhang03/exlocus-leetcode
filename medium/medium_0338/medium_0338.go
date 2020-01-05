package medium_0338

func countOne(num int) int {
	result := 0
	for num != 0 {
		tmp := (num & 1)
		if tmp == 1 {
			result++
		}
		num = num >> 1
	}
	return result
}

func countBits(num int) []int {
	result := []int{}
	if num >= 0 {
		for i := 0; i <= num; i++ {
			result = append(result, countOne(i))
		}
	}

	return result
}
