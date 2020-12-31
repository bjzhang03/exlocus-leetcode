package easy_0257

import (
	"strconv"
)

func deepfirstsearch(root *TreeNode, path string, result *[]string) {
	if root != nil {
		if root.Left == nil && root.Right == nil {
			tmp := (path + strconv.Itoa(root.Val))
			*result = append(*result, tmp)
		}
		if root.Left != nil {
			tmp := path + strconv.Itoa(root.Val) + "->"
			deepfirstsearch(root.Left, tmp, result)
		}

		if root.Right != nil {
			tmp := path + strconv.Itoa(root.Val) + "->"
			deepfirstsearch(root.Right, tmp, result)
		}

	}
}

func binaryTreePaths(root *TreeNode) []string {
	if root != nil {
		result := []string{}
		path := ""
		deepfirstsearch(root, path, &result)
		return result
	}
	return nil
}
