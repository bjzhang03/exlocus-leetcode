package medium_0508

import "math"

func findFrequentTreeSum(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		used := map[int]int{}
		treeSum(root, &used)
		// 找到map中的最大值
		mcount := math.MinInt32
		for _, val := range used {
			if mcount < val {
				mcount = val
			}
		}
		// 将出现次数最多的数据拿出来
		for key, val := range used {
			if mcount == val {
				result = append(result, key)
			}
		}
	}
	return result
}

// 求tree的sum,并把sum的数据记录在map中
func treeSum(root *TreeNode, used *map[int]int) int {
	if root != nil {
		result := 0
		if root.Left == nil && root.Right == nil {
			result = result + root.Val
		} else if root.Left != nil && root.Right == nil {
			result = result + root.Val + treeSum(root.Left, used)
		} else if root.Left == nil && root.Right != nil {
			result = result + root.Val + treeSum(root.Right, used)
		} else if root.Left != nil && root.Right != nil {
			result = result + root.Val + treeSum(root.Left, used) + treeSum(root.Right, used)
		}
		if _, ok := (*used)[result]; ok {
			(*used)[result]++
		} else if !ok {
			(*used)[result] = 1
		}
		return result
	}
	return 0
}
