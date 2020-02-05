package medium_0450

// bst二叉搜索树,不是二叉平衡树
// 二叉搜索出不需要左右旋转
func deleteNode(root *TreeNode, key int) *TreeNode {
	// 找到需要删除的节点
	if root != nil {
		if root.Val == key {
			// 判断当前数据是否正确
			//叶子节点
			if root.Left == nil && root.Right == nil {
				return nil
			} else if root.Left == nil && root.Right != nil {
				root.Val = bstMin(root.Right)
				root.Right = deleteNode(root.Right, root.Val)
			} else if root.Left != nil && root.Right == nil {
				root.Val = bstMax(root.Left)
				root.Left = deleteNode(root.Left, root.Val)
			} else if root.Left != nil && root.Right != nil {
				root.Val = bstMax(root.Left)
				root.Left = deleteNode(root.Left, root.Val)
			}
			return root
		} else if root.Val > key {
			root.Left = deleteNode(root.Left, key)
		} else if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		}
	}
	return root
}

//获取bst中的最大值
func bstMax(root *TreeNode) int {
	result := root.Val
	if root.Right != nil {
		result = bstMax(root.Right)
	}
	return result
}

// 获取bst中的最小值
func bstMin(root *TreeNode) int {
	result := root.Val
	if root.Left != nil {
		result = bstMin(root.Left)
	}
	return result
}
