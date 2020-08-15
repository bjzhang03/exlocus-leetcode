package easy_0083

import (
	"fmt"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	var dd = []struct {
		head     *ListNode
		expected *ListNode
	}{
		{buildLists([]int{1, 1, 2}), buildLists([]int{1, 2})},
		{buildLists([]int{1, 1, 2, 3, 3}), buildLists([]int{1, 2, 3})},
	}

	for _, val := range dd {
		actual := deleteDuplicates(val.head)

		if !listsEqual(val.expected, actual) {
			t.Errorf("Test Failed! expected := %s, actual := %s", listSprint(val.expected), listSprint(actual))
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

/*打印list的函数*/
func listSprint(l *ListNode) string {
	result := ""
	for l != nil {
		result = result + fmt.Sprint(l.Val)
		if l.Next != nil {
			result = result + "->"
		}
		l = l.Next
	}
	return result
}
