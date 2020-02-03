package medium_0092

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	result := &ListNode{0, nil}
	if head != nil {
		count := 1
		item := head
		tmp := result
		// 前面部分的数据直接复制过来
		for count < m {
			tmp.Next = &ListNode{item.Val, nil}
			item = item.Next
			tmp = tmp.Next
			count++
		}
		// 中间部分的数据需要进行逆转
		reverse := &ListNode{0, nil}

		for count <= n {
			tmpsave := &ListNode{item.Val, reverse.Next}
			item = item.Next
			// 链表的头插法
			reverse.Next = tmpsave
			count++
		}

		//fmt.Println("count = ", count)
		tmprev := reverse.Next
		for tmprev.Next != nil {
			tmprev = tmprev.Next

		}
		for item != nil {
			//fmt.Println(item.Val)
			tmprev.Next = &ListNode{item.Val, nil}
			item = item.Next
			tmprev = tmprev.Next
		}
		// 拼接起来
		tmp.Next = reverse.Next
	}
	return result.Next
}
