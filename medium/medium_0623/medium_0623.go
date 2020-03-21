package medium_0623

func addOneRow(root *TreeNode, v int, d int) *TreeNode {
	if root != nil && d == 1 {
		// 处理特殊情况的数据
		result := &TreeNode{v, root, nil}
		return result
	} else if root != nil && d > 0 {
		level := 1
		nodes := []*TreeNode{root}
		for len(nodes) > 0 {
			// 找到特定的层级,直接添加数据进来即可
			if level == d-1 {
				for i := 0; i < len(nodes); i++ {
					tLeft := &TreeNode{v, nodes[i].Left, nil}
					tRight := &TreeNode{v, nil, nodes[i].Right}

					nodes[i].Left = tLeft
					nodes[i].Right = tRight
				}
				break
			} else {
				// 层次遍历操作
				nnodes := []*TreeNode{}
				for i := 0; i < len(nodes); i++ {
					if nodes[i].Left != nil {
						nnodes = append(nnodes, nodes[i].Left)
					}
					if nodes[i].Right != nil {
						nnodes = append(nnodes, nodes[i].Right)
					}
				}
				// 更新node节点的数据
				nodes = nodes[0:0]
				nodes = append(nodes, nnodes...)
				level++
			}
		}
	}
	return root
}
