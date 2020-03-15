package medium_0567

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > 0 && len(s2) > 0 {
		s1m := map[uint8]int{}
		for i := 0; i < len(s1); i++ {
			if _, ok := s1m[s1[i]]; ok {
				s1m[s1[i]]++
			} else if !ok {
				s1m[s1[i]] = 1
			}
		}

		s2sta, s2end := 0, 0
		s2m := map[uint8]int{}
		for s2sta < len(s2) && s2end < len(s2) {
			// 然后对数据进行判断操作
			if maplow(s1m, s2m) {
				s2m[s2[s2sta]]--
				// 如果为0,则删除
				if s2m[s2[s2sta]] == 0 {
					delete(s2m, s2[s2sta])
				}
				s2sta++
			} else {
				if _, ok := s2m[s2[s2end]]; ok {
					s2m[s2[s2end]]++
				} else if !ok {
					s2m[s2[s2end]] = 1
				}
				s2end++
			}

			// 直接决定数字是否满足
			if mapequal(s1m, s2m) {
				return true
			}
		}
		// 处理特殊情况s2end到了,但是s2sta没到
		for s2sta < len(s2) {
			if maplow(s1m, s2m) {
				s2m[s2[s2sta]]--
				// 如果为0,则删除
				if s2m[s2[s2sta]] == 0 {
					delete(s2m, s2[s2sta])
				}
			}
			// 直接决定数字是否满足
			if mapequal(s1m, s2m) {
				return true
			}
			s2sta++
		}
	}
	return false
}

// map 相等
func mapequal(a, b map[uint8]int) bool {
	return func(am, bm map[uint8]int) bool {
		// a==b
		for key, val := range am {
			if _, ok := bm[key]; !(ok && b[key] == val) {
				return false
			}
		}
		for key, val := range bm {
			if _, ok := am[key]; !(ok && am[key] == val) {
				return false
			}
		}
		return true
	}(a, b)
}

func maplow(a, b map[uint8]int) bool {
	return func(am, bm map[uint8]int) bool {
		// a<b
		for key, val := range am {
			if _, ok := bm[key]; !(ok && b[key] >= val) {
				return false
			}
		}
		return true
	}(a, b)
}
