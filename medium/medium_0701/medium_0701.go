package medium_0701

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val > val {
			if root.Left == nil {
				root.Left = &TreeNode{val, nil, nil}
			} else {
				insertIntoBST(root.Left, val)
			}
		} else if root.Val < val {
			if root.Right == nil {
				root.Right = &TreeNode{val, nil, nil}
			} else {
				insertIntoBST(root.Right, val)
			}
		}
	}
	return root
}
