package medium_0438

func findAnagrams(s string, p string) []int {
	result := []int{}
	if len(s) > 0 && len(p) > 0 {
		// 统计p中数字出现的频次
		pmap := make(map[uint8]int, 26)
		smap := make(map[uint8]int, 26)
		for i := 0; i < len(p); i++ {
			if _, ok := pmap[p[i]]; !ok {
				pmap[p[i]] = 1
			} else if ok {
				pmap[p[i]]++
			}
		}
		// 统计到当前位置的数据出现的频次
		mapslice := make([]map[uint8]int, len(s))
		for i := 0; i < len(s); i++ {
			if _, ok := smap[s[i]]; !ok {
				smap[s[i]] = 1
			} else if ok {
				smap[s[i]]++
			}
			mapslice[i] = map[uint8]int{}
			for key, val := range smap {
				mapslice[i][key] = val
			}
		}
		// 挨个查找是否正确
		for i := 0; i < len(s)-len(p)+1; i++ {
			cmap := mapSub(mapslice[i+len(p)-1], mapslice[i])
			//检查当前字符在不在map中
			if _, ok := cmap[s[i]]; !ok {
				cmap[s[i]] = 1
			} else if ok {
				cmap[s[i]]++
			}
			// 判断map是否相等
			if mapEquals(pmap, cmap) {
				result = append(result, i)
			}
		}
	}
	return result
}

func mapEquals(amap map[uint8]int, bmap map[uint8]int) bool {
	result := true
	for key, val := range amap {
		if _, ok := bmap[key]; !(ok && val == bmap[key]) {
			result = false
		}
	}
	return result
}

func mapSub(amap map[uint8]int, bmap map[uint8]int) map[uint8]int {
	result := make(map[uint8]int, 26)
	// 默认amap是比较大的数据,得到的是amap-bmap
	for key, val := range amap {
		if _, ok := bmap[key]; ok && val-bmap[key] > 0 {
			result[key] = val - bmap[key]
		} else if !ok && key > 0 {
			result[key] = val
		}
	}
	return result
}
