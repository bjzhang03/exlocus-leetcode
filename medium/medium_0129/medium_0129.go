package medium_0129
<<<<<<< HEAD:medium/medium_0129/medium_0129.go
=======

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
>>>>>>> ccf62001e728049c9fa47645ff1f406b58578a46:medium/medium_0129/medium_0129.go

func dfs(root *TreeNode, current int, save *int) {
	if root.Left == nil && root.Right == nil {
		current = current*10 + root.Val
		*save = *save + current
		current = current / 10
	} else {

		if root.Left != nil {
			current = current*10 + root.Val
			dfs(root.Left, current, save)
			current = current / 10
		}
		if root.Right != nil {
			current = current*10 + root.Val
			dfs(root.Right, current, save)
			current = current / 10
		}
	}

}

func sumNumbers(root *TreeNode) int {
	result := 0
	if root != nil {
		dfs(root, 0, &result)
	}
	return result
}
