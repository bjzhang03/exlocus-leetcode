package medium_0151

// 不用处理s长度为0的情况
func reverse(s string) string {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result = result + string(s[i])
	}
	return result
}

// 压缩中间出现的多个空白数据
func compressSpace(s string) string {
	if len(s) > 0 {
		result := ""
		for i := 0; i < len(s); i++ {
			for s[i] == ' ' && i+1 < len(s) && s[i+1] == ' ' {
				i++
			}
			result = result + string(s[i])
		}
		return result
	}
	return s
}

func reverseWords(s string) string {
	// 去掉前面的space
	for len(s) > 0 && s[0] == ' ' {
		s = s[1:]
	}
	// 去掉最后面的space
	for len(s) > 0 && s[len(s)-1] == ' ' {
		s = s[:len(s)-1]
	}
	// 中间多个空白字符需要压缩
	s = compressSpace(s)

	if len(s) > 0 {
		// 先将整个字符串reverse
		s = reverse(s)
		// 再将word进行reverse
		result := ""
		ista := 0
		iend := 0
		for iend < len(s) {
			if s[iend] == ' ' {
				tmp := s[ista:iend]
				result = result + reverse(tmp)
				ista = iend + 1
				result = result + " "
			}
			iend++
		}
		// 加上最后的数据
		tmp := s[ista:iend]
		result = result + reverse(tmp)
		return result
	}
	return s
}
