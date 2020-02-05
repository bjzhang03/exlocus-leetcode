package medium_0452

import "sort"

// 这个题目我没看明白,也不会做,参考的网上的思路
// https://www.cnblogs.com/joelwang/p/10636934.html
func findMinArrowShots(points [][]int) int {
	result := 0
	if len(points) > 0 {
		// 先对数据进行排序操作
		save := &intervals{points}
		sort.Sort(save)

		for len(points) > 0 {
			current := points[0][1]
			// 先找到当前的最短的数据
			for i := 0; i < len(points); i++ {
				if current > points[i][1] {
					current = points[i][1]
				}
			}
			// 移除可以删掉的数据
			npoints := [][]int{}
			for i := 0; i < len(points); i++ {
				if !func(p int, point []int) bool {
					if point[0] <= p && point[1] >= p {
						return true
					}
					return false
				}(current, points[i]) {
					npoints = append(npoints, points[i])
				}
			}
			points = append(points[0:0], npoints...)
			result++
		}

	}

	return result

}

type intervals struct {
	data [][]int
}

func (in intervals) Len() int {
	return len(in.data)
}

func (in intervals) Less(i, j int) bool {
	if in.data[i][0] < in.data[j][0] {
		return true
	} else if in.data[i][0] == in.data[j][0] && in.data[i][1] < in.data[j][1] {
		return true
	}
	return false
}

func (in intervals) Swap(i, j int) {
	tmp := append([]int{}, in.data[i]...)
	in.data[i] = in.data[i][0:0]
	in.data[i] = append(in.data[i], in.data[j]...)
	in.data[j] = in.data[j][0:0]
	in.data[j] = append(in.data[j], tmp...)
}
