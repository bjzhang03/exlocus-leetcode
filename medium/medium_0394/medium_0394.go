package medium_0394

import (
	"strconv"
)

func decodeString(s string) string {
	result := ""
	if len(s) > 0 {
		str := s
		for containRight(str) {
			eni := 0
			sti := 0
			// 找到结束的地方
			for i := 0; i < len(str); i++ {
				if str[i] == ']' {
					eni = i
					break
				}
			}
			// 找到开始的地方
			for i := eni; i >= 0; i-- {
				if str[i] == '[' {
					// 开始向前面寻找数字
					for i > 0 {
						i--
						if !(str[i] <= '9' && str[i] >= '0') {
							break
						}
					}
					// 如果不是第一个的话,则需要后移一位
					if i > 0 {
						i++
					}
					// 找到进行标记,然后再退出即可
					sti = i
					break
				}
			}

			str = str[:sti] + decodeItem(str[sti:eni+1]) + str[eni+1:]
		}
		result = str
	}
	return result
}

func containRight(str string) bool {
	for i := 0; i < len(str); i++ {
		if str[i] == ']' {
			return true
		}
	}
	return false
}

func decodeItem(str string) string {
	result := ""
	// 寻找数字
	sti := 0
	eni := 0
	for eni < len(str) {
		if str[eni] == '[' {
			break
		}
		eni++
	}
	count, _ := strconv.Atoi(str[sti:eni])
	// 寻找字母
	sti = eni + 1
	for eni < len(str) {
		if str[eni] == ']' {
			break
		}
		eni++
	}
	cstr := str[sti:eni]
	// 拼接
	for i := 0; i < count; i++ {
		result = result + cstr
	}
	return result
}
