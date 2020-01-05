package medium_0337

//https://blog.csdn.net/happyaaaaaaaaaaa/article/details/50880121
func robHelper(root *TreeNode) (int, int) {
	robr, robc := 0, 0
	if root != nil {
		leftr, leftc := robHelper(root.Left)
		rightr, rightc := robHelper(root.Right)
		// 记录下rob到孩子节点的数据
		robc = leftr + rightr
		// 记录下rob到自己的数据
		if leftc+rightc+root.Val > robc {
			robr = leftc + rightc + root.Val
		} else {
			robr = robc
		}
	}
	return robr, robc
}

func rob(root *TreeNode) int {
	result := 0
	if root != nil {
		result, _ = robHelper(root)
	}
	return result
}
