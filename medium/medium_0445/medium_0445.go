package medium_0445

// 链表相加
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 != nil && l2 == nil {
		return l1
	} else if l1 == nil && l2 != nil {
		return l2
	} else if l1 != nil && l2 != nil {
		// 使用堆栈的思路来进行处理数据
		sl1 := []int{}
		sl2 := []int{}
		for l1 != nil {
			sl1 = append(sl1, l1.Val)
			l1 = l1.Next
		}
		for l2 != nil {
			sl2 = append(sl2, l2.Val)
			l2 = l2.Next
		}

		result := &ListNode{0, nil}
		current := 0
		for len(sl1) > 0 && len(sl2) > 0 {
			tmp := &ListNode{(sl2[len(sl2)-1] + sl1[len(sl1)-1] + current) % 10, nil}
			current = (sl2[len(sl2)-1] + sl1[len(sl1)-1] + current) / 10
			tmp.Next = result.Next
			result.Next = tmp
			// 退栈
			sl1 = sl1[:len(sl1)-1]
			sl2 = sl2[:len(sl2)-1]
		}

		for len(sl1) > 0 {
			tmp := &ListNode{(sl1[len(sl1)-1] + current) % 10, nil}
			current = (sl1[len(sl1)-1] + current) / 10
			tmp.Next = result.Next
			result.Next = tmp
			sl1 = sl1[:len(sl1)-1]
		}

		for len(sl2) > 0 {
			tmp := &ListNode{(sl2[len(sl2)-1] + current) % 10, nil}
			current = (sl2[len(sl2)-1] + current) / 10
			tmp.Next = result.Next
			result.Next = tmp
			sl2 = sl2[:len(sl2)-1]
		}

		if current > 0 {
			tmp := &ListNode{current, nil}
			tmp.Next = result.Next
			result.Next = tmp
		}

		return result.Next
	}
	return nil
}
