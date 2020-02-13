package medium_0474

import "strconv"

// 这个题目是非常经典的0-1背包问题
func findMaxForm(strs []string, m int, n int) int {
	return solve(strs, len(strs)-1, m, n, &map[string]int{})
}

func solve(strs []string, si, mi, ni int, used *map[string]int) int {
	result := 0
	if _, ok := (*used)[statusToString(si, mi, ni)]; ok {
		result = (*used)[statusToString(si, mi, ni)]

	} else if si >= 0 && si < len(strs) && mi >= 0 && ni >= 0 {
		result = solve(strs, si-1, mi, ni, used)
		tmpm, tmpn := subStr(mi, ni, strs[si])
		if tmpm >= 0 && tmpn >= 0 && result < 1+solve(strs, si-1, tmpm, tmpn, used) {
			result = 1 + solve(strs, si-1, tmpm, tmpn, used)
		}
	}
	(*used)[statusToString(si, mi, ni)] = result
	return result
}

// golang的map不支持使用自定义的struct作为key,所以这里将struct转化为string
func statusToString(si, mi, ni int) string {
	return "(" + strconv.Itoa(si) + "," + strconv.Itoa(mi) + "," + strconv.Itoa(ni) + ")"
}

func subStr(m, n int, str string) (int, int) {
	zeroc, onec := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '0' {
			zeroc++
		} else if str[i] == '1' {
			onec++
		}
	}
	return m - zeroc, n - onec
}
