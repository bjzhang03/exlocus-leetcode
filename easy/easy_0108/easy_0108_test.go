package easy_0108

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	var satbst = []struct {
		nums     []int
		expected []int
	}{
		{[]int{-10, -3, 0, 5, 9}, []int{-10, -3, 0, 5, 9}},
		{[]int{-10, -3, 1, 2, 5, 9}, []int{-10, -3, 1, 2, 5, 9}},
	}

	for _, val := range satbst {
		actual := sortedArrayToBST(val.nums)
		/*对数据进行排序,最后的结果应该是一致的*/
		sort.Ints(val.expected)

		if !reflect.DeepEqual(inorderVisit(actual), val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(inorderVisit(actual)))
		}
	}
}

/*中序遍历的操作,二叉搜索树的中序遍历应该刚好是有序的结果*/
func inorderVisit(root *TreeNode) []int {
	if root != nil {
		result := append([]int{}, inorderVisit(root.Left)...)
		result = append(result, root.Val)
		result = append(result, inorderVisit(root.Right)...)
		return result
	}
	return []int{}
}
