package medium_0648

import (
	"sort"
	"strings"
)

func replaceWords(dict []string, sentence string) string {
	words := strings.Split(sentence, " ")
	if len(words) > 0 && len(dict) > 0 {
		// 先对数据进行排序操作
		handleDict := &dicts{dict}
		sort.Sort(handleDict)

		// 找到最短的字符前缀进行替换即可
		for i := 0; i < len(words); i++ {
			for j := 0; j < len(dict); j++ {
				if len(dict[j]) > len(words[i]) {
					break
				} else {
					if isprefix(dict[j], words[i]) {
						words[i] = dict[j]
						break
					}
				}
			}
		}
		result := ""
		for i := 0; i < len(words); i++ {
			result = result + words[i] + " "
		}
		return result[:len(result)-1]
	}
	return sentence
}

func isprefix(a, b string) bool {
	if len(a) <= len(b) {
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	return false
}

type dicts struct {
	data []string
}

func (d dicts) Len() int {
	return len(d.data)
}

func (d dicts) Less(i, j int) bool {
	if len(d.data[i]) < len(d.data[j]) {
		return true
	} else if len(d.data[i]) == len(d.data[j]) {
		for k := 0; k < len(d.data[i]); k++ {
			if d.data[i][k] < d.data[j][k] {
				return true
			} else if d.data[i][k] > d.data[j][k] {
				return false
			}
		}
	}
	return false
}

func (d dicts) Swap(i, j int) {
	tmp := d.data[i]
	d.data[i] = d.data[j]
	d.data[j] = tmp
}
