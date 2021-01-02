package medium_0050

import (
	"fmt"
)

func myPow(x float64, n int) float64 {
	return func(candidate float64, pow int) float64 {
		if pow > 0 {
			return solve(candidate, pow, map[string]float64{})
		} else if pow < 0 {
			return 1.0 / solve(candidate, 0-pow, map[string]float64{})
		}
		return 1.0
	}(x, n)
}

func solve(x float64, n int, save map[string]float64) float64 {
	//如果数据存在,则直接返回回去即可
	if val, ok := save[fmt.Sprintf("%s,%s", fmt.Sprint(x), fmt.Sprint(n))]; ok {
		return val
	} else if !ok {
		result := 1.0
		if n == 1 {
			result = result * x
		} else if n > 1 {
			if n%2 != 0 {
				result = result * x * solve(x, n/2, save) * solve(x, n/2, save)
			} else if n%2 == 0 {
				result = result * solve(x, n/2, save) * solve(x, n/2, save)
			}
		}
		save[fmt.Sprintf("%s,%s", fmt.Sprint(x), fmt.Sprint(n))] = result
		return result
	}
	return 0
}
