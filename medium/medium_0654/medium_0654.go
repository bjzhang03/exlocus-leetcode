package medium_0654

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) > 0 {
		// 找到数字最大的数据的,和它的index
		currentMax := nums[0]
		currentIndex := 0
		for i := 1; i < len(nums); i++ {
			if nums[i] > currentMax {
				currentMax = nums[i]
				currentIndex = i
			}
		}
		// 分割左右子树
		left := nums[:currentIndex]
		right := nums[currentIndex+1:]
		// 新建数据
		root := &TreeNode{currentMax, nil, nil}
		//分别构建左右子树
		if len(left) > 0 {
			root.Left = constructMaximumBinaryTree(left)
		}
		if len(right) > 0 {
			root.Right = constructMaximumBinaryTree(right)
		}
		return root
	}
	return nil
}
