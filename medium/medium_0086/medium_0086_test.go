package medium_0086

import (
	"fmt"
	"testing"
)

func TestPartition(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		head     *ListNode
		x        int
		expected *ListNode
	}{
		{buildLists([]int{1, 4, 3, 2, 5, 2}), 3, buildLists([]int{1, 2, 2, 4, 3, 5})},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := partition(val.head, val.x)
		/*验证结果*/
		if !listsEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}

/*根据数组构建lists*/
func buildLists(vals []int) *ListNode {
	header := &ListNode{0, nil}
	pointer := header

	for _, val := range vals {
		tmp := &ListNode{val, nil}
		pointer.Next = tmp
		pointer = pointer.Next
	}
	return header.Next
}

/*判断两个list是否相等的函数*/
func listsEqual(f *ListNode, s *ListNode) bool {
	if f == nil && s == nil {
		return true
	} else if f != nil && s != nil {
		for f != nil && s != nil {
			if f.Val != s.Val {
				return false
			}
			if f.Next == nil && s.Next != nil || f.Next != nil && s.Next == nil {
				return false
			}
			f = f.Next
			s = s.Next
		}
		return true
	}
	return false
}
