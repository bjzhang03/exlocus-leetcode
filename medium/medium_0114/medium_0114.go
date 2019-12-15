package medium_0114

func flattenTree(root *TreeNode) *TreeNode {
	if root != nil {
		// 取出左子树
		left := root.Left
		// 取出右子树
		right := root.Right

		root.Left = nil
		root.Right = flattenTree(left)
		// 右子树下游分支
		nextRight := root

		for nextRight.Right != nil {
			nextRight = nextRight.Right
		}
		nextRight.Right = flattenTree(right)
	}
	return root
}

func flatten(root *TreeNode) {
	flattenTree(root)
}
