package medium_0424

// 滑动窗口的问题,这种问题我以前没有见过,第一次处理
// https://www.cnblogs.com/migoo/p/12227706.html
func characterReplacement(s string, k int) int {
	result := 0
	if len(s) > 0 {
		// 计算所有的数字出现的频次
		save := make(map[uint8]int, len(s))
		csave := make([]map[uint8]int, len(s))
		for i := 0; i < len(s); i++ {
			if _, ok := save[s[i]]; !ok {
				save[s[i]] = 1
			} else {
				save[s[i]]++
			}
			// 将当前的map存储下来
			csave[i] = map[uint8]int{}
			for key, val := range save {
				csave[i][key] = val
			}
		}

		sti, eni := 0, 0
		for eni < len(csave) {
			for eni < len(csave) && removeKchars(s, sti, eni, k, csave) {
				if result < (eni - sti + 1) {
					result = (eni - sti + 1)
				}
				eni++
			}
			if eni < len(csave) && !removeKchars(s, sti, eni, k, csave) {
				eni++
				sti++
			}
		}
	}
	return result
}

// 删除k个字符可以变成同一种字符测试
func removeKchars(s string, sti, eni, k int, csave []map[uint8]int) bool {
	// 字符串长度
	result := false
	slen := eni - sti + 1
	for key, val := range csave[eni] {
		// 判断第一个字符串的长度
		if key == s[sti] {
			if (val - csave[sti][key] + 1 + k) >= slen {
				result = true
			}
		} else {
			if _, ok := csave[sti][key]; ok && (k+val-csave[sti][key]) >= slen {
				result = true
			} else if !ok && (val+k) >= slen {
				result = true
			}
		}
	}
	return result
}
