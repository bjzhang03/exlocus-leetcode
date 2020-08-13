package medium_0652

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	if root != nil {
		result := []*TreeNode{}
		solve(root, &[]*TreeNode{}, &result)
		return result
	}
	return nil
}

func solve(root *TreeNode, used *[]*TreeNode, result *[]*TreeNode) {
	if root != nil {
		for i := 0; i < len(*used); i++ {
			// 此时出现一样的数据了
			if isSame(root, (*used)[i]) {
				// 查看以前是否已经处理过了,没有添加到result中的话则添加到result中去
				flag := true
				for j := 0; j < len(*result); j++ {
					if isSame(root, (*result)[j]) {
						flag = false
						break
					}
				}
				if flag {
					(*result) = append(*result, root)
				}
			}
		}
		(*used) = append(*used, root)
		// 再分别求解左右子树的数据
		solve(root.Left, used, result)
		solve(root.Right, used, result)
	}
}

// 判断两棵二叉树是否是一样的
func isSame(a, b *TreeNode) bool {
	if a == nil && b == nil {
		return true
	} else if a != nil && b != nil {
		if a.Val == b.Val {
			return isSame(a.Left, b.Left) && isSame(a.Right, b.Right)
		}
	}
	return false
}
