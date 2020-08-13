package medium_0467

// 这个题目实在是没看懂到底该怎么做
func findSubstringInWraproundString(p string) int {
	result := 0
	if len(p) > 0 {
		used := map[string]bool{}
		for i := 0; i < len(p); i++ {
			//fmt.Println(used)
			eni := i + 1
			for ; eni < len(p); eni++ {
				if _, ok := used[p[i:eni]]; !ok && checkWrap(p[i:eni], used) {
					used[p[i:eni]] = true
				} else if !checkWrap(p[i:eni], used) {
					break
				}
			}
			// 处理一直到最后一个字符的问题
			if func() bool {
				if eni == len(p) {
					if _, ok := used[p[i:]];
						!ok && checkWrap(p[i:], used) {
						return true
					}
				}
				return false
			}() {
				used[p[i:]] = true
			}
		}
		//fmt.Println(len(used), used)
		return len(used)
	}
	return result
}

// 检查字符串是否是连续的字符串
func checkWrap(s string, used map[string]bool) bool {
	if _, ok := used[s]; ok {
		return true
	} else if len(s) == 1 {
		return true
	} else if len(s) > 1 {
		if s[len(s)-1] != (s[len(s)-2]-'a'+1)%26+'a' {
			return false
		} else {
			return true
		}
	}
	return false
}
