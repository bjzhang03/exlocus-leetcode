package medium_0451

import "sort"

func frequencySort(s string) string {
	if len(s) > 0 {
		// 数据预处理
		// 统计数据
		sfmap := map[uint8]int{}
		for i := 0; i < len(s); i++ {
			if _, ok := sfmap[s[i]]; !ok {
				sfmap[s[i]] = 1
			} else if ok {
				sfmap[s[i]]++
			}
		}
		// 将数据一致的放到一起来处理
		subStr := []string{}
		for key, val := range sfmap {
			tmp := ""
			for i := 0; i < val; i++ {
				tmp = tmp + string(key)
			}
			subStr = append(subStr, tmp)
		}
		// 对数据进行排序
		save := &sortf{subStr, sfmap}
		sort.Sort(save)
		// 重组数据
		result := ""
		for i := len(subStr) - 1; i >= 0; i-- {
			result = result + subStr[i]
		}
		return result
	}
	return ""
}

type sortf struct {
	suint []string
	sfmap map[uint8]int
}

func (sf sortf) Len() int {
	return len(sf.suint)
}
func (sf sortf) Less(i, j int) bool {
	return sf.sfmap[sf.suint[i][0]] < sf.sfmap[sf.suint[j][0]]
}

func (sf sortf) Swap(i, j int) {
	tmp := sf.suint[i]
	sf.suint[i] = sf.suint[j]
	sf.suint[j] = tmp
}
