package easy_0203

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) String() string {
	result := ""
	if l != nil {
		result = "["
		tmp := l
		for tmp != nil && tmp.Next != nil {
			result = result + fmt.Sprint(tmp.Val) + ","
			tmp = tmp.Next
		}
		result = result + fmt.Sprint(tmp.Val) + "]"
	}
	return result
}
