package easy_0107

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLevelOrderBottom(t *testing.T) {
	var lob = []struct {
		root     *TreeNode
		expected [][]int
	}{
		{buildTreeFromSlice([]int{3, 9, 20, -1, -1, 15, 7}, 0), [][]int{{15, 7}, {9, 20}, {3}}},
	}

	for _, val := range lob {
		actual := levelOrderBottom(val.root)

		if !reflect.DeepEqual(actual, val.expected) {
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
