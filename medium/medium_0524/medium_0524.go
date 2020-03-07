package medium_0524

import (
	"sort"
)

func findLongestWord(s string, d []string) string {
	if len(s) > 0 && len(d) > 0 {
		// 首先对d进行排序
		saved := &strsort{d}
		sort.Sort(saved)
		for i := 0; i < len(d); i++ {
			if func(a, b string) bool {
				// 判断当前的数据是否可以通过删除字符获得
				asta := 0
				used := ""
				for i := 0; i < len(b); i++ {
					for asta < len(a) {
						if a[asta] != b[i] {
							asta++
						} else {
							used = used + string(a[asta])
							asta++
							break
						}
					}
				}
				return len(used) == len(b)
			}(s, d[i]) {
				return d[i]
			}
		}
	}
	return ""
}

// 自定义排序操作
type strsort struct {
	data []string
}

func (in strsort) Len() int {
	return len(in.data)
}

func (in strsort) Less(i, j int) bool {
	if len(in.data[i]) > len(in.data[j]) {
		return true
	} else if len(in.data[i]) == len(in.data[j]) {
		return func(a, b string) bool {
			for i := 0; i < len(a); i++ {
				if a[i] < b[i] {
					return true
				} else if a[i] > b[i] {
					return false
				}
			}
			return false
		}(in.data[i], in.data[j])
	}
	return false
}

func (in strsort) Swap(i, j int) {
	tmp := in.data[i]
	in.data[i] = in.data[j]
	in.data[j] = tmp
}
