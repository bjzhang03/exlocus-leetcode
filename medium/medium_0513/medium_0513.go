package medium_0513

func findBottomLeftValue(root *TreeNode) int {
	result := 0
	if root != nil {
		current := []*TreeNode{root}
		for len(current) > 0 {
			// 记下下一次loop的数据
			nloop := []*TreeNode{}
			for i := 0; i < len(current); i++ {
				if current[i].Left != nil {
					nloop = append(nloop, current[i].Left)
				}
				if current[i].Right != nil {
					nloop = append(nloop, current[i].Right)
				}
			}
			// 判断下一层是不是已经没有数据了
			if len(nloop) == 0 {
				result = current[0].Val
			}
			// 更新下一次循环的数据
			current = current[0:0]
			current = append(current, nloop...)
		}
	}
	return result
}
