package medium_0515

import "math"

func largestValues(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		current := []*TreeNode{root}
		for len(current) > 0 {
			nloop := []*TreeNode{}
			tmp := math.MinInt32
			for i := 0; i < len(current); i++ {
				if tmp < current[i].Val {
					tmp = current[i].Val
				}
				if current[i].Left != nil {
					nloop = append(nloop, current[i].Left)
				}
				if current[i].Right != nil {
					nloop = append(nloop, current[i].Right)
				}
			}
			result = append(result, tmp)
			current = current[0:0]
			current = append(current, nloop...)
		}
	}
	return result
}
