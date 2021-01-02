package medium_0056

import (
	"sort"
)

func merge(intervals [][]int) [][]int {
	// 这里只需要进行一次排序即可,减少时间的消耗
	h := &handler{intervals: intervals}
	sort.Sort(h)
	return solve(h)
}

func solve(h *handler) [][]int {
	// 先对h进行排序
	//sort.Sort(h)
	//fmt.Printf("%s\n", fmt.Sprint(h.intervals))
	for i := 0; i < h.Len(); i++ {
		// 判断当前的数据是否是可以合并的
		if i+1 < h.Len() && func(id1, id2 int, save [][]int) bool {
			if save[id1][1] >= save[id2][0] {
				return true
			}
			return false
		}(i, i+1, h.intervals) {
			// 出现可以合并的则直接合并
			if h.intervals[i][1] < h.intervals[i+1][1] {
				return solve(&handler{intervals: append(append(append([][]int{}, h.intervals[:i]...), []int{h.intervals[i][0], h.intervals[i+1][1]}), h.intervals[i+2:]...)})
			} else if h.intervals[i][1] >= h.intervals[i+1][1] {
				return solve(&handler{intervals: append(append(append([][]int{}, h.intervals[:i]...), []int{h.intervals[i][0], h.intervals[i][1]}), h.intervals[i+2:]...)})
			}
		}
	}
	// 没有出现可以合并的数据,则直接返回回来即可
	return h.intervals
}

/*辅助处理的数据结构*/
type handler struct {
	intervals [][]int
}

// 获取此slice的长度
func (h handler) Len() int { return len(h.intervals) }

// 自定义的排序算法
func (h handler) Less(i, j int) bool {
	if h.intervals[i][0] < h.intervals[j][0] {
		return true
	}
	return false
}

// 交换数据
func (h handler) Swap(i, j int) {
	tmp := append([]int{}, h.intervals[i]...)
	h.intervals[i] = append(h.intervals[i][0:0], h.intervals[j]...)
	h.intervals[j] = append(h.intervals[j][0:0], tmp...)
}
