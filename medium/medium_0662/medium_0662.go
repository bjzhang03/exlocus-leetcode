package medium_0662

// https://blog.csdn.net/jackzhang_123/article/details/78866977
// 这个题目我没有想明白是咋解决,参考了上面的方案
func widthOfBinaryTree(root *TreeNode) int {
	if root != nil {

		result := 0
		rootwn := deepFirstSearch(root, 1)
		current := []*TreeNodewn{rootwn}
		for len(current) > 0 {
			if result < current[len(current)-1].number-current[0].number+1 {
				result = current[len(current)-1].number - current[0].number + 1
			}
			// 开始下一层次的循环
			ncurrent := []*TreeNodewn{}
			for i := 0; i < len(current); i++ {
				if current[i].Left != nil {
					ncurrent = append(ncurrent, current[i].Left)
				}
				if current[i].Right != nil {
					ncurrent = append(ncurrent, current[i].Right)
				}
			}
			current = current[:0]
			current = append(current, ncurrent...)
		}
		return result
	}
	return 0
}

// 构建一个带有标记的treenode数据结构
type TreeNodewn struct {
	Val    int
	Left   *TreeNodewn
	Right  *TreeNodewn
	number int
}

// 先深度优先搜索构造数列
func deepFirstSearch(root *TreeNode, number int) *TreeNodewn {
	if root != nil {
		result := &TreeNodewn{root.Val, nil, nil, number}
		if root.Left != nil {
			result.Left = deepFirstSearch(root.Left, number*2)
		}
		if root.Right != nil {
			result.Right = deepFirstSearch(root.Right, number*2+1)
		}
		return result
	}
	return nil
}
