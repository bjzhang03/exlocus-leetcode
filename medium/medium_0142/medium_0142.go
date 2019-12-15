package medium_0142

import "fmt"

func detectCycle(head *ListNode) *ListNode {
	// 判断有没有环,无环则返回nil
	slow := head
	fast := head

	// 快慢指针判断有没有环型
	for fast != nil && fast.Next != nil {
		// 到这里发现fast多走了一个环的距离
		if slow == fast {
			// 有环型,则找到环开始的地方
			fmt.Println(fast.Val)
			break
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast == nil || fast.Next == nil {
		return nil
	}

	// 快指针刚好走了一个环的距离
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	fmt.Println(fast.Val)
	return fast

}
