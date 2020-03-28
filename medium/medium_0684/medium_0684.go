package medium_0684

func findRedundantConnection(edges [][]int) []int {
	if len(edges) > 0 {
		for i := 0; i < len(edges); i++ {
			if breadthFirstSearch(edges[:i], edges[i]) {
				return append([]int{}, edges[i]...)
			}
		}
	}
	return nil
}

// 广度优先算法来寻找路径
// 这里存在优化的方案,就是可以每次记住当前已经使用过的数据,来进行优化
// 这样就不需要每次都是从头开始全部查找了
func breadthFirstSearch(edges [][]int, target []int) bool {
	if len(edges) > 0 && len(target) > 0 {
		save := []int{target[0]}
		used := map[int]bool{target[0]: true}
		for len(save) > 0 {
			current := save[0]
			// 已经使用过的数字
			for j := 0; j < len(edges); j++ {
				if _, ok1 := used[edges[j][1]]; !ok1 && edges[j][0] == current {
					used[edges[j][1]] = true
					save = append(save, edges[j][1])
				}
				if _, ok2 := used[edges[j][0]]; !ok2 && edges[j][1] == current {
					used[edges[j][0]] = true
					save = append(save, edges[j][0])
				}
			}
			save = save[1:]
			if _, ok := used[target[1]]; ok {
				return true
			}
		}
	}
	return false
}
