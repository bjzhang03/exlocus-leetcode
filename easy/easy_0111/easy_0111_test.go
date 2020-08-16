package easy_0111

import (
	"fmt"
	"math"
	"testing"
)

func TestMinDepth(t *testing.T) {
	var md = []struct {
		root     *TreeNode
		expected int
	}{
		{buildTreeFromSlice([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7}, 0), 2},
	}

	for _, val := range md {
		actual := minDepth(val.root)

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
