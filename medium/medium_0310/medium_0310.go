package medium_0310

import "fmt"

func solve(edges [][]int, result *[]int) {

	fmt.Println("edges = ", edges, result)

	nodesta := map[int]int{}
	nodeend := map[int]int{}

	for _, val := range edges {
		// 计算第一个点出现的次数
		if _, ok := nodesta[val[0]]; !ok {
			nodesta[val[0]] = 1
		} else if ok {
			nodesta[val[0]]++
		}
		// 计算第二个点出现的次数
		if _, ok := nodeend[val[1]]; !ok {
			nodeend[val[1]] = 1
		} else if ok {
			nodeend[val[1]]++
		}
	}

	// 去掉所有的那些边缘的边
	nextedges := [][]int{}
	for _, val := range edges {
		tmp0 := 0
		tmp1 := 0
		if _, ok := nodesta[val[0]]; ok && nodesta[val[0]] == 1 {
			tmp0++
		}

		if _, ok := nodeend[val[0]]; ok && nodeend[val[0]] == 1 {
			tmp0++
		}

		if _, ok := nodesta[val[1]]; ok && nodesta[val[1]] == 1 {
			tmp1++
		}
		if _, ok := nodeend[val[1]]; ok && nodeend[val[1]] == 1 {
			tmp1++
		}

		if !(tmp0 == 1 || tmp1 == 1) {
			nextedges = append(nextedges, val)
		}
	}

	if len(nextedges) == 0 {
		// 全部点已经用完,找到那个大于1的点即可
		for key, val := range nodesta {
			if val > 1 {
				(*result) = append(*result, key)
				return
			}
		}
		for key, val := range nodeend {
			if val > 1 {
				(*result) = append(*result, key)
				return
			}
		}
	} else if len(nextedges) == 1 {
		*result = append(*result, nextedges[0]...)
		return
	} else if len(nextedges) == 2 {
		count := map[int]int{}
		for i := 0; i < len(nextedges); i++ {
			for j := 0; j < len(nextedges[0]); j++ {
				if _, ok := count[nextedges[i][j]]; !ok {
					count[nextedges[i][j]] = 1
				} else if ok {
					count[nextedges[i][j]]++
				}
			}
		}

		for key, val := range count {
			if val > 1 {
				(*result) = append(*result, key)
				return
			}
		}
	} else if len(nextedges) > 1 {
		solve(nextedges, result)
	}

}

func findMinHeightTrees(n int, edges [][]int) []int {
	result := []int{}
	if n > 0 && len(edges) == 1 {
		result = append(result, edges[0]...)
	} else if n > 0 && len(edges) > 0 {
		solve(edges, &result)
	} else if n > 0 {
		for i := 0; i < n; i++ {
			result = append(result, i)
		}
	}
	return result
}
