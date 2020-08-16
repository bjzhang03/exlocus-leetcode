package easy_0110

import (
	"fmt"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	var ib = []struct {
		root     *TreeNode
		expected bool
	}{
		{buildTreeFromSlice([]int{3, 9, 20, -1, -1, 15, 7}, 0), true},
		{buildTreeFromSlice([]int{1, 2, 2, 3, 3, -1, -1, 4, 4}, 0), false},
	}

	for _, val := range ib {
		actual := isBalanced(val.root)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}

/*根据slice来构造一个二叉树*/
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
