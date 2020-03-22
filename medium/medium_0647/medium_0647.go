package medium_0647

import "fmt"

func countSubstrings(s string) int {
	if len(s) > 0 {
		result := 0
		used := map[string]bool{}
		for i := 0; i < len(s); i++ {
			for j := i; j >= 0; j-- {
				if ispalindromic(s, []int{j, i}, &used) {
					result++
				}
			}
		}
		return result
	}
	return 0
}

func ispalindromic(s string, indexs []int, used *map[string]bool) bool {
	if indexs[0] >= indexs[1] {
		return true
	} else if _, ok := (*used)[fmt.Sprint(indexs)]; ok {
		return (*used)[fmt.Sprint(indexs)]
	} else if s[indexs[0]] == s[indexs[1]] && ispalindromic(s, []int{indexs[0] + 1, indexs[1] - 1}, used) {
		(*used)[fmt.Sprint(indexs)] = true
		return true
	}
	(*used)[fmt.Sprint(indexs)] = false
	return false
}
