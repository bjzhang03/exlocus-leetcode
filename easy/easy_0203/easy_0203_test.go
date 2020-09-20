package easy_0203

import (
	"fmt"
	"testing"
)

func TestRemoveElements(t *testing.T) {

	var cases = []struct {
		head     *ListNode
		val      int
		expected *ListNode
	}{
		{buildListFromSlice([]int{1, 2, 6, 3, 4, 5, 6}), 6, buildListFromSlice([]int{1, 2, 3, 4, 5})},
	}

	for _, val := range cases {
		actual := removeElements(val.head, val.val)

		if !isListEqual(val.expected, actual) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}

// 从slice来构建list列表
func buildListFromSlice(nums []int) *ListNode {
	header := &ListNode{0, nil}
	tmp := header

	for i := 0; i < len(nums); i++ {
		tmp.Next = &ListNode{nums[i], nil}
		tmp = tmp.Next
	}
	return header.Next
}

// 判断两个List是否相等
func isListEqual(one *ListNode, two *ListNode) bool {
	for one != nil && two != nil {
		if one.Val != two.Val {
			return false
		}
		one = one.Next
		two = two.Next
	}
	if one != nil || two != nil {
		return false
	}
	return true
}
