package medium_0785

func isBipartite(graph [][]int) bool {
	return solve(graph)
}

func solve(graph [][]int) bool {
	if len(graph) > 0 {
		// 用染色法来判断数据是不是二部图
		save := []int{}
		used := map[int]bool{}
		white := map[int]bool{} // 一开始的节点染成白色的
		black := map[int]bool{}
		// 判断是否所有的节点都已经使用过了
		for func(save *[]int, used *map[int]bool, white *map[int]bool, black *map[int]bool, graph *[][]int) bool {
			// 如果save为0的时候,则找到第一个不是空的数据
			if len(*save) == 0 {
				for i := 0; i < len(*graph); i++ {
					if _, ok := (*used)[i]; !ok && len((*graph)[i]) > 0 {
						*save = append(*save, i)
						(*used)[i] = true
						(*white)[i] = true
						return true
					}
				}
			}
			return len((*save)) > 0
		}(&save, &used, &white, &black, &graph) {
			// 拿出当前的需要处理的点
			current := save[0]
			save = save[1:]
			// 当前的点是白色的
			if _, okw := white[current]; okw {
				// 所有的于这个点相连的点都只能是黑色的
				for i := 0; i < len(graph[current]); i++ {
					if _, oke := white[graph[current][i]]; oke {
						return false
					} else if !oke {
						black[graph[current][i]] = true
					}
					if _, oku := used[graph[current][i]]; !oku {
						save = append(save, graph[current][i])
						used[graph[current][i]] = true
					}
				}
			} else if _, okb := black[current]; okb { // 当前的点是黑色的
				// 所有的于这个点相连的点都只能是黑色的
				for i := 0; i < len(graph[current]); i++ {
					if _, oke := black[graph[current][i]]; oke {
						return false
					} else if !oke {
						white[graph[current][i]] = true
					}
					if _, oku := used[graph[current][i]]; !oku {
						save = append(save, graph[current][i])
						used[graph[current][i]] = true
					}
				}
			}
		}
		return true
	}
	return false
}
