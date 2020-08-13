package medium_0095
<<<<<<< HEAD:medium/medium_0095/medium_0095.go
=======

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
>>>>>>> ccf62001e728049c9fa47645ff1f406b58578a46:medium/medium_0095/medium_0095.go

// 递归的思想进行处理
func generateTreesRange(save []int, ins int, ine int) []*TreeNode {
	result := []*TreeNode{}
	if ins < ine {
		for i := ins; i < ine; i++ {
			//fmt.Println("i = ", i)
			left := generateTreesRange(save, ins, i)
			right := generateTreesRange(save, i+1, ine)
			//fmt.Println(left)
			//fmt.Println(right)
			for j := 0; j < len(left); j++ {
				for k := 0; k < len(right); k++ {
					root := &TreeNode{save[i], nil, nil}
					root.Left = left[j]
					root.Right = right[k]
					result = append(result, root)
				}
			}
		}
	} else {
		// 添加一个nil进来
		result = append(result, nil)
	}
	return result

}

func generateTrees(n int) []*TreeNode {
	result := []*TreeNode{}
	if n > 0 {
		save := []int{}
		for i := 0; i < n; i++ {
			save = append(save, i+1)
		}
		result = generateTreesRange(save, 0, len(save))
	}
	return result
}
