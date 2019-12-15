package medium_0236

func commonAncestor(root, p, q *TreeNode) (*TreeNode, bool, bool) {
	if root != nil {
		left, leftp, leftq := commonAncestor(root.Left, p, q)
		right, rightp, rightq := commonAncestor(root.Right, p, q)
		// 对当前的情况的处理
		if root == p && !(leftq || rightq) {
			return root, true, false
		} else if root == q && !(leftp || rightp) {
			return root, false, true
		} else if root == p && (leftq || rightq) {
			return root, true, true
		} else if root == q && (leftp || rightp) {
			return root, true, true
		} else if leftp && leftq {
			return left, true, true
		} else if rightp && rightq {
			return right, true, true
		} else if (leftp || rightp) && (leftq || rightq) {
			return root, true, true
		} else if !(leftp || rightp) && (leftq || rightq) {
			return root, false, true
		} else if (leftp || rightp) && !(leftq || rightq) {
			return root, true, false
		}
	}
	return root, false, false
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	result, _, _ := commonAncestor(root, p, q)
	return result
}
