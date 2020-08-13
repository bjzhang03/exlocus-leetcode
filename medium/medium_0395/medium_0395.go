package medium_0395

import (
	"sort"
)

func longestSubstring(s string, k int) int {
	result := 0
	// 确保s的长度大于0,并且k不能大于s
	if len(s) > 0 && k <= len(s) {
		// 先统计所有的数据各自出现的次数
		savm := map[uint8][]int{}
		save := map[uint8]int{}
		for i := 0; i < len(s); i++ {
			if _, ok := save[s[i]]; !ok {
				save[s[i]] = 1
				savm[s[i]] = []int{i}

			} else {
				save[s[i]]++
				savm[s[i]] = append(savm[s[i]], i)
			}
		}
		// 寻找分割点
		split := []int{}
		for key, val := range save {
			if val < k {
				split = append(split, savm[key]...)
			}
		}
		// 没有分割点的话,则直接返回即可
		if len(split) == 0 {
			result = len(s)
		} else {
			// 对分割点下边进行排序
			sort.Ints(split)
			sti := 0
			for i := 0; i < len(split); i++ {
				tmp := longestSubstring(s[sti:split[i]], k)
				if result < tmp {
					result = tmp
				}
				if split[i]+1 < len(s) {
					sti = split[i] + 1
				}
			}
			// 计算出现最后的数据是最长数据的情况
			if sti != 0 {
				tmp := longestSubstring(s[sti:], k)
				if result < tmp {
					result = tmp
				}
			}
		}
	}
	return result
}
