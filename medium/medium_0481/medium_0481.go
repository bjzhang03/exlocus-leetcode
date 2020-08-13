package medium_0481

func magicalString(n int) int {
	result := 0
	if n == 1 || n == 2 || n == 3 {
		result = 1
	} else if n > 3 {
		save := make([]int, n+5)
		save[1] = 1
		save[2] = 2
		save[3] = 2
		sti := 3
		eni := 3

		for eni <= n {
			if save[eni] == 2 {
				for i := 0; i < save[sti]; i++ {
					eni++
					save[eni] = 1
				}
				sti++
			} else if save[eni] == 1 {
				for i := 0; i < save[sti]; i++ {
					eni++
					save[eni] = 2
				}
				sti++
			}
		}
		// 统计1的个数
		for i := 0; i <= n; i++ {
			if save[i] == 1 {
				result++
			}
		}
	}
	return result
}
