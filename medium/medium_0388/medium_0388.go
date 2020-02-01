package medium_0388

import (
	"strings"
)

func lengthLongestPath(input string) int {
	result := 0
	if len(input) > 0 {
		strs := strings.Split(input, "\n")
		save := make([][2]int, len(strs))

		for i := 0; i < len(strs); i++ {
			// 计算前面有多少个\t
			current := getCountT(strs[i])
			// 计算当前的长度
			count := getCharCout(strs[i])

			for j := i - 1; j >= 0; j-- {
				if save[j][0]+1 == current {
					count = save[j][1] + count
					break
				}
			}
			save[i][0] = current
			save[i][1] = count

			if result < count+current && isFile(strs[i]) {
				result = count + current
			}
		}
	}
	return result
}

// 判断字符串前面有多少个\t
func getCountT(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\t' {
			result++
		} else {
			break
		}
	}
	return result
}

// 获取出去\t之外的长度
func getCharCout(str string) int {
	result := 0
	for i := 0; i < len(str); i++ {
		if str[i] != '\t' {
			result++
		}
	}
	return result
}

// 判断是否是file
func isFile(str string) bool {
	result := false

	for i := 0; i < len(str); i++ {
		if str[i] == '.' {
			return true
		}
	}

	return result
}
