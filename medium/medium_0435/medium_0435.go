package medium_0435

import (
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	result := 0
	if len(intervals) > 0 {
		// 先对数据进行排序
		sortinters := &inters{intervals}
		sort.Sort(sortinters)
		save := make([][][]int, len(intervals))
		for i := 0; i < len(intervals); i++ {
			tmps := [][]int{intervals[i]}
			for j := i - 1; j >= 0; j-- {
				if checkIntervals(save[j], intervals[i]) && len(tmps) < (len(save[j])+1) {
					tmps = append([][]int{intervals[i]}, save[j]...)
				}
			}
			save[i] = tmps
			// 找到可以添加的最多的数据
			if result < len(tmps) {
				result = len(tmps)
			}
		}
		// 最后取反即可
		result = len(intervals) - result
	}
	return result
}

func checkIntervals(intervals [][]int, item []int) bool {
	// intervals 里面的数据已经排序过了,所以直接看第一个即可
	result := true
	if len(intervals) > 0 && !(intervals[0][0] < item[0] && intervals[0][1] <= item[0]) {
		result = false
	}
	return result
}

type inters struct {
	data [][]int
}

// 获取此slice的长度
func (in inters) Len() int { return len(in.data) }

// 自定义的排序算法
func (in inters) Less(i, j int) bool {
	if in.data[i][0] < in.data[j][0] {
		return true
	} else if in.data[i][0] == in.data[j][0] && in.data[i][1] < in.data[j][1] {
		return true
	}
	return false
}

// 交换数据
func (in inters) Swap(i, j int) {
	tmp := append([]int{}, in.data[i]...)
	in.data[i] = in.data[i][0:0]
	in.data[i] = append(in.data[i], in.data[j]...)
	in.data[j] = in.data[j][0:0]
	in.data[j] = append(in.data[j], tmp...)
}
