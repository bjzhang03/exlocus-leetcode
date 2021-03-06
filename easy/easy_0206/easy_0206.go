package easy_0206

func reverseList(head *ListNode) *ListNode {
	if head != nil {
		result := &ListNode{head.Val, nil}
		head = head.Next
		for head != nil {
			tmp := &ListNode{head.Val, nil}
			tmp.Next = result
			result = tmp
			head = head.Next
		}
		return result
	}
	return nil
}
