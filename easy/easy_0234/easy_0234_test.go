package easy_0234

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	var cases = []struct {
		head     *ListNode
		expected bool
	}{
		{buildLists([]int{}), true},
		{buildLists([]int{1, 2}), false},
		{buildLists([]int{1, 2, 2, 1}), true},
	}

	for _, val := range cases {
		actual := isPalindrome(val.head)

		if actual != val.expected {
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
