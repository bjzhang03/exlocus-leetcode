package medium_0148

func splitList(head *ListNode) (*ListNode, *ListNode) {
	// 如果是nil的话则直接返回
	if head == nil {
		return nil, nil
	} else if head != nil && head.Next == nil {
		left := head
		right := head.Next
		head.Next = nil
		return left, right
	} else {
		// 快慢指针切分链表
		slow := head
		fast := head

		// 使用right记录数据,需要清空最后left的next的数据
		right := slow
		for fast != nil && fast.Next != nil {
			right = slow
			slow = slow.Next
			fast = fast.Next.Next
		}

		left := head
		// 清空left的最后的next的数据
		right.Next = nil
		return left, slow
	}

	return nil, nil
}

func mergeList(left *ListNode, right *ListNode) *ListNode {
	// 临时节点
	root := &ListNode{0, nil}
	tmp := root

	for left != nil && right != nil {
		if left.Val <= right.Val {
			tmp.Next = left
			left = left.Next
			tmp = tmp.Next
			tmp.Next = nil
		} else {
			tmp.Next = right
			right = right.Next
			tmp = tmp.Next
			tmp.Next = nil
		}
	}
	// 如果还有节点的话则直接添加到后面即可
	if left != nil {
		tmp.Next = left
	}
	if right != nil {
		tmp.Next = right
	}

	return root.Next
}

func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	} else if head != nil && head.Next == nil {
		return head
	} else {
		left, right := splitList(head)

		return mergeList(sortList(left), sortList(right))
	}
	return nil
}
