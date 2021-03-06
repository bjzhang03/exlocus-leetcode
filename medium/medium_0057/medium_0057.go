package medium_0057

import "sort"

func insert(intervals [][]int, newInterval []int) [][]int {
	h := &handler{intervals: append(intervals, newInterval)}
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
func (in handler) Len() int { return len(in.intervals) }

// 自定义的排序算法
func (in handler) Less(i, j int) bool {
	if in.intervals[i][0] < in.intervals[j][0] {
		return true
	}
	return false
}

// 交换数据
func (in handler) Swap(i, j int) {
	tmp := append([]int{}, in.intervals[i]...)
	in.intervals[i] = append(in.intervals[i][0:0], in.intervals[j]...)
	in.intervals[j] = append(in.intervals[j][0:0], tmp...)
}
