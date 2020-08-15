package easy_0101

import (
	"fmt"
	"testing"
)

func TestIsSymmetric(t *testing.T) {
	var sy = []struct {
		root     *TreeNode
		expected bool
	}{
		{buildTreeFromSlice([]int{1, 2, 2, 3, 4, 4, 3}, 0), true},
		{buildTreeFromSlice([]int{1, 2, 2, -1, 3, -1, 3}, 0), false},
	}

	for _, val := range sy {
		actual := isSymmetric(val.root)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}

func buildTreeFromSlice(nums []int, position int) *TreeNode {
	/*这里使用-1来约定数据是nil的*/
	if position >= 0 && position < len(nums) && nums[position] != -1 {
		root := &TreeNode{nums[position], nil, nil}
		root.Left = buildTreeFromSlice(nums, position*2+1)
		root.Right = buildTreeFromSlice(nums, position*2+2)
		return root
	}
	return nil
}
