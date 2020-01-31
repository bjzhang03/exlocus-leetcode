package medium_0373

import "sort"

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	save := pairs{}

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			save.data = append(save.data, pair{nums1[i], nums2[j]})
		}
	}
	sort.Sort(save)

	result := [][]int{}

	for i := 0; i < k && i < len(save.data); i++ {
		result = append(result, []int{save.data[i].x, save.data[i].y})
	}
	return result
}

type pair struct {
	x int
	y int
}

type pairs struct {
	data []pair
}

// 获取此slice的长度
func (p pairs) Len() int { return len(p.data) }

// 自定义的排序算法
func (p pairs) Less(i, j int) bool {
	return p.data[i].x+p.data[i].y < p.data[j].x+p.data[j].y
}

// 交换数据
func (p pairs) Swap(i, j int) {
	tmp := &pair{p.data[i].x, p.data[i].y}
	p.data[i].x = p.data[j].x
	p.data[i].y = p.data[j].y
	p.data[j].x = tmp.x
	p.data[j].y = tmp.y
}
