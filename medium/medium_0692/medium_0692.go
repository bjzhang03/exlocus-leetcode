package medium_0692

import "sort"

func topKFrequent(words []string, k int) []string {
	if len(words) > 0 && k > 0 {
		// 先统计words中每一个单词出现的次数
		save := map[string]int{}
		for i := 0; i < len(words); i++ {
			if _, ok := save[words[i]]; ok {
				save[words[i]]++
			} else if !ok {
				save[words[i]] = 1
			}
		}
		data := []string{}
		for key, _ := range save {
			data = append(data, key)
		}
		sort.Sort(&frequent{data, save})
		if len(data) <= k {
			return data
		} else if len(data) > k {
			return data[:k]
		}
	}
	return nil
}

// 这里我自己懒得写堆排序了,所以直接调用排序接口进行排序了
type frequent struct {
	data []string
	save map[string]int
}

func (f frequent) Len() int {
	return len(f.data)
}

func (f frequent) Less(i, j int) bool {
	if f.save[f.data[i]] < f.save[f.data[j]] {
		return false
	} else if f.save[f.data[i]] == f.save[f.data[j]] {
		for k := 0; k < len(f.data[i]); k++ {
			if f.data[i] < f.data[j] {
				return true
			} else if f.data[i] > f.data[j] {
				return false
			}
		}
	}
	return true
}

func (f frequent) Swap(i, j int) {
	tmp := f.data[i]
	f.data[i] = f.data[j]
	f.data[j] = tmp
}
