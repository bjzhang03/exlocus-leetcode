package medium_0002

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//新建结果数据
	result := &ListNode{0, nil}
	it := result
	carry := 0

	for ; l1 != nil && l2 != nil; {
		//计算当前的值
		temp := &ListNode{(l1.Val + l2.Val + carry) % 10, nil}
		it.Next = temp
		//计算进位数据
		carry = (l1.Val + l2.Val + carry) / 10
		//后移一位数据
		it = it.Next
		l1 = l1.Next
		l2 = l2.Next
	}
	//l1没有到nil
	for ; l1 != nil; {
		//计算当前的值
		temp := &ListNode{(l1.Val + carry) % 10, nil}
		it.Next = temp
		//计算进位数据
		carry = (l1.Val + carry) / 10
		//后移一位数据
		it = it.Next
		l1 = l1.Next
	}
	//l2没有到nil
	for ; l2 != nil; {
		//计算当前的值
		temp := &ListNode{(l2.Val + carry) % 10, nil}
		it.Next = temp
		//计算进位数据
		carry = (l2.Val + carry) / 10
		//后移一位数据
		it = it.Next
		l2 = l2.Next
	}
	//carry还有数据
	if carry > 0 {
		temp := &ListNode{carry, nil}
		it.Next = temp
		//计算进位数据
	}

	return result.Next
}
