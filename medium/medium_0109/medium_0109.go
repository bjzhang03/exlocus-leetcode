package medium_0109

func buildTreeFromSlice(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		result := &TreeNode{nums[0], nil, nil}
		return result
	} else {
		povt := len(nums) / 2
		result := &TreeNode{nums[povt], nil, nil}
		result.Left = buildTreeFromSlice(nums[:povt])
		result.Right = buildTreeFromSlice(nums[povt+1:])
		return result
	}
	return nil
}

func sortedListToBST(head *ListNode) *TreeNode {
	// 先获取所有的数据
	nums := []int{}
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return buildTreeFromSlice(nums)
}
