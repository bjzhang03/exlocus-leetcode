package medium_0576

func findPaths(m int, n int, N int, i int, j int) int {
	if m > 0 && n > 0 && N > 0 {
		MOD := 1000000007
		result := 0
		current := map[int]int{pointToKey(i, j): 1}
		for N > 0 {
			ncurrent := map[int]int{}
			for key, val := range current {
				x, y := keyToPoint(key)
				//上下左右判断
				if pass(m, n, x-1, y) {
					if _, ok := ncurrent[pointToKey(x-1, y)]; ok {
						ncurrent[pointToKey(x-1, y)] = (ncurrent[pointToKey(x-1, y)] + val) % MOD
					} else if !ok {
						ncurrent[pointToKey(x-1, y)] = val % MOD
					}
				} else {
					result = (result + val) % MOD
				}
				if pass(m, n, x+1, y) {
					if _, ok := ncurrent[pointToKey(x+1, y)]; ok {
						ncurrent[pointToKey(x+1, y)] = (ncurrent[pointToKey(x+1, y)] + val) % MOD
					} else if !ok {
						ncurrent[pointToKey(x+1, y)] = val % MOD
					}
				} else {
					result = (result + val) % MOD
				}
				if pass(m, n, x, y-1) {
					if _, ok := ncurrent[pointToKey(x, y-1)]; ok {
						ncurrent[pointToKey(x, y-1)] = (ncurrent[pointToKey(x, y-1)] + val) % MOD
					} else if !ok {
						ncurrent[pointToKey(x, y-1)] = val % MOD
					}
				} else {
					result = (result + val) % MOD
				}
				if pass(m, n, x, y+1) {
					if _, ok := ncurrent[pointToKey(x, y+1)]; ok {
						ncurrent[pointToKey(x, y+1)] = (ncurrent[pointToKey(x, y+1)] + val) % MOD
					} else if !ok {
						ncurrent[pointToKey(x, y+1)] = val % MOD
					}
				} else {
					result = (result + val) % MOD
				}
			}
			//更新current
			current = map[int]int{}
			for key, val := range ncurrent {
				current[key] = val
			}
			N--
		}
		return result
	}
	return 0
}

func keyToPoint(key int) (int, int) {
	return key % 1000, key / 1000
}

func pointToKey(x, y int) int {
	return x + y*1000
}

func pass(m, n int, x, y int) bool {
	if x >= 0 && x < m && y >= 0 && y < n {
		return true
	}
	return false
}
