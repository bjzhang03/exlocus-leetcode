package easy_0083

func deleteDuplicates(head *ListNode) *ListNode {
	// 新建一个节点
	result := &ListNode{0, nil}
	if head != nil {
		// 找到节点的起始位置
		it := result
		for ; head != nil; {
			it.Next = &ListNode{head.Val, nil}
			it = it.Next
			// 去掉重复的数据
			for ; head.Next != nil && head.Val == head.Next.Val; {
				head = head.Next
			}
			if head.Next == nil && head.Val == it.Val {
				head = head.Next
			} else {
				// 后移一步
				head = head.Next
			}
		}
	}
	return result.Next
}
