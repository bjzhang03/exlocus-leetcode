package medium_0725

func splitListToParts(root *ListNode, k int) []*ListNode {
	result := []*ListNode{}
	// 先计算出链表的长度
	llen := 0
	item := root
	for item != nil {
		llen++
		item = item.Next
	}
	// 计算相除结果和余数
	dres := llen / k
	drem := llen % k
	// 拼接链表
	for k > 0 {
		// 使用临时变量
		tmproot := &ListNode{0, nil}
		tmpitem := tmproot
		// 根据结果数字进行构建
		for i := 0; i < dres; i++ {
			tmpitem.Next = &ListNode{root.Val, nil}
			tmpitem = tmpitem.Next
			root = root.Next
		}
		// 根据余数进行构建
		if drem > 0 {
			tmpitem.Next = &ListNode{root.Val, nil}
			tmpitem.Next = tmpitem.Next
			root = root.Next
			drem--
		}
		k--
		result = append(result, tmproot.Next)
	}
	return result
}
