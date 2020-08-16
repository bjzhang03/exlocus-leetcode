package easy_0112

import (
	"fmt"
	"math"
	"testing"
)

func TestHasPathSum(t *testing.T) {
	var hps = []struct {
		root     *TreeNode
		sum      int
		expected bool
	}{
		{buildTreeFromSlice([]int{5, 4, 8, 11, math.MinInt32, 13, 4, 7, 2, math.MinInt32, math.MinInt32, math.MinInt32, 1}, 0), 22, true},
	}

	for _, val := range hps {
		actual := hasPathSum(val.root, val.sum)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}

/*根据slice来构造一个二叉树*/
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
