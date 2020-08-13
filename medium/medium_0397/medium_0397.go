package medium_0397

// 使用全局变量便于重复使用
// 因为我知道这里的包被多次使用,因此偷懒
var save = map[int]int{}

func integerReplacement(n int) int {
	result := 0
	if n > 1 {
		if _, ok := save[n]; ok {
			result = save[n]
		} else if _, ok := save[n]; !ok {
			if n%2 == 0 {
				result = integerReplacement(n/2) + 1
			} else {
				add := integerReplacement(n+1) + 1
				sub := integerReplacement(n-1) + 1
				if add < sub {
					result = add
				} else {
					result = sub
				}
				save[n] = result
			}
		}
	}
	return result
}
