package easy_0100

import (
	"fmt"
	"math"
	"testing"
)

func TestIsSameTree(t *testing.T) {

	var is = []struct {
		p        *TreeNode
		q        *TreeNode
		expected bool
	}{
		{buildTreeFromSlice([]int{1, 2, 3}, 0), buildTreeFromSlice([]int{1, 2, 3}, 0), true},
		{buildTreeFromSlice([]int{1, 2}, 0), buildTreeFromSlice([]int{1, math.MinInt32, 2}, 0), false},
		{buildTreeFromSlice([]int{1, 2, 1}, 0), buildTreeFromSlice([]int{1, 1, 2}, 0), false},
	}

	for _, val := range is {
		actual := isSameTree(val.p, val.q)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

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
