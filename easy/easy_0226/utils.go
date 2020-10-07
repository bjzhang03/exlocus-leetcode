package easy_0226

import (
	"fmt"
	"math"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTreeFromSlice(nums []int, position int) *TreeNode {
	/*这里使用math.MinInt32来约定数据是nil的*/
	if position >= 0 && position < len(nums) && nums[position] != math.MinInt32 {
		root := &TreeNode{nums[position], nil, nil}
		root.Left = buildTreeFromSlice(nums, position*2+1)
		root.Right = buildTreeFromSlice(nums, position*2+2)
		return root
	}
	return nil
}

func (t *TreeNode) String() string {
	return fmt.Sprint(inorderer(t))
}

func (t *TreeNode) Equal(s *TreeNode) bool {
	if t == nil && s == nil {
		return true
	} else if t != nil && s != nil {
		if t.Val != s.Val {
			return false
		} else {
			return t.Left.Equal(s.Left) && t.Right.Equal(s.Right)
		}
	}
	return false
}

func inorderer(t *TreeNode) []int {
	result := []int{}
	if t != nil {
		result = append(result, t.Val)
		result = append(result, inorderer(t.Left)...)
		result = append(result, inorderer(t.Right)...)
	}
	return result
}
