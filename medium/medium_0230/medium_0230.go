package medium_0230
<<<<<<< HEAD:medium/medium_0230/medium_0230.go
=======

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
>>>>>>> ccf62001e728049c9fa47645ff1f406b58578a46:medium/medium_0230/medium_0230.go

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
