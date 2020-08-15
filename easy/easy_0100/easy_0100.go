package easy_0100

func isSameTree(p *TreeNode, q *TreeNode) bool {
	result := true
	if p != nil && q == nil {
		result = false
	} else if p == nil && q != nil {
		result = false
	} else if p != nil && q != nil && p.Val != q.Val {
		result = false
	} else if p == nil && q == nil {
		result = true
	} else if p != nil && q != nil && p.Val == q.Val {
		result = result && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
	}
	return result
}
