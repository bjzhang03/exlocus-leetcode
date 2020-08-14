package medium_0230

func getKnums(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		left := getKnums(root.Left)
		result = append(result, left...)
		result = append(result, root.Val)
		result = append(result, getKnums(root.Right)...)
	}
	return result
}

func kthSmallest(root *TreeNode, k int) int {
	return getKnums(root)[k-1]
}
