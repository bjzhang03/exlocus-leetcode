package medium_0646

import "sort"

func findLongestChain(pairs [][]int) int {
	if len(pairs) > 0 {
		chainData := &chains{pairs}
		sort.Sort(chainData)
		// 向前查找数据
		save := make([][][]int, len(pairs))
		result := 0
		for i := 0; i < len(pairs); i++ {
			tmp := [][]int{pairs[i]}
			for j := i - 1; j >= 0; j-- {
				if save[j][len(save[j])-1][1] < pairs[i][0] && len(save[j])+1 > len(tmp) {
					tmp = append(append([][]int{}, save[j]...), pairs[i])
				}
			}
			save[i] = tmp
			if result < len(tmp) {
				result = len(tmp)
			}
		}
		return result
	}
	return 0
}

// 构建数据结构,便于实现排序操作
type chains struct {
	pair [][]int
}

func (c chains) Len() int {
	return len(c.pair)
}

func (c chains) Less(i, j int) bool {
	if c.pair[i][0] < c.pair[j][0] {
		return true
	} else if c.pair[i][0] == c.pair[j][0] && c.pair[i][1] < c.pair[j][1] {
		return true
	}
	return false
}

func (c chains) Swap(i, j int) {
	tmp := c.pair[i]
	c.pair[i] = append([]int{}, c.pair[j]...)
	c.pair[j] = append([]int{}, tmp...)
}
