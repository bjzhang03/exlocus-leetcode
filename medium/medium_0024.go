package medium

func swapPairs(head *ListNode) *ListNode {
	if head != nil {
		result := &ListNode{0, nil}
		item := head
		p := result
		for item != nil && item.Next != nil {
			p.Next = &ListNode{item.Next.Val, nil}
			p = p.Next
			p.Next = &ListNode{item.Val, nil}
			p = p.Next
			item = item.Next.Next
		}
		if item != nil {
			p.Next = &ListNode{item.Val, nil}
		}
		return result.Next
	}
	return nil
}
