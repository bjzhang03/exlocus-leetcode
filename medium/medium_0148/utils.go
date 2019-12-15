package medium_0148

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListFromSlice(nums []int) *ListNode {

	root := &ListNode{0, nil}

	tmp := root

	for _, key := range nums {
		tmp.Next = &ListNode{key, nil}
		tmp = tmp.Next
	}

	return root.Next

}
