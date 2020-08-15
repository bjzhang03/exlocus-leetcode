package easy_0104

func maxDepth(root *TreeNode) int {
	result := 0
	if root != nil {
		leftDepth := maxDepth(root.Left)
		rightDepth := maxDepth(root.Right)
		if leftDepth > rightDepth {
			result = leftDepth + 1
		} else {
			result = rightDepth + 1
		}
	}
	return result
}
