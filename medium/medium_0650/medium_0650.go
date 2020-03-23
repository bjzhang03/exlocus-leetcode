package medium_0650

func minSteps(n int) int {
	if n > 0 {
		save := make([]int, n+1)
		save[1] = 1
		for i := 2; i <= n; i++ {
			tmp := 10000000000
			for j := i - 1; j >= 1; j-- {
				if i%j == 0 && tmp > i/j+save[j] {
					tmp = i/j + save[j]
				}
			}
			save[i] = tmp
		}
		return save[n] - 1
	}
	return 0
}
