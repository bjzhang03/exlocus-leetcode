package easy_0257

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestBinaryTreePaths(t *testing.T) {

	var cases = []struct {
		root     *TreeNode
		expected []string
	}{
		{buildTreeFromSlice([]int{1, 2, 3, math.MinInt32, 5}, 0), []string{"1->2->5", "1->3"}},
	}

	for _, val := range cases {
		actual := binaryTreePaths(val.root)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
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
