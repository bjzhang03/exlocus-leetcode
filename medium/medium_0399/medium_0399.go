package medium_0399

import "fmt"

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	result := []float64{}
	if len(queries) > 0 {
		// 先将string映射成数据下标
		save := map[string]int{}
		count := 0
		for i := 0; i < len(equations); i++ {
			for j := 0; j < len(equations[i]); j++ {
				if _, ok := save[equations[i][j]]; !ok {
					save[equations[i][j]] = count
					count++
				}
			}
		}

		// 根据下标构建乘法表
		multable := make([][]float64, len(save))
		for i := 0; i < len(multable); i++ {
			multable[i] = make([]float64, len(save))
		}
		// 初始化所有的数据为-1
		for i := 0; i < len(multable); i++ {
			for j := 0; j < len(multable[i]); j++ {
				multable[i][j] = -1
			}
		}

		// 先将已经存在的数据填写进去
		for i := 0; i < len(equations); i++ {
			multable[save[equations[i][0]]][save[equations[i][1]]] = values[i]
		}

		for i := 0; i < len(multable); i++ {
			multable[i][i] = 1
			for j := i + 1; j < len(multable[i]); j++ {
				if multable[i][j] != -1 {
					multable[j][i] = 1.0 / multable[i][j]
				}
				if multable[j][i] != -1 {
					multable[i][j] = 1.0 / multable[j][i]
				}
			}
		}

		for i := 0; i < len(queries); i++ {
			_, ok1 := save[queries[i][0]]
			_, ok2 := save[queries[i][1]]

			if ok1 && ok2 {
				_, re := deepFirstSearch(&multable, save[queries[i][0]], save[queries[i][1]], &map[int]bool{})
				result = append(result, re)
			} else {
				result = append(result, -1)
			}
		}

		fmt.Println(result)
	}
	return result
}

// 深度优先遍历,并记下路径上的数据
func deepFirstSearch(multable *[][]float64, sti, eni int, used *map[int]bool) (bool, float64) {

	fmt.Println("sti := ", sti, " ,eni := ", eni)
	result := float64(-1)
	flag := false
	if (*multable)[sti][eni] != -1 {
		flag = true
		result = (*multable)[sti][eni]
	} else if (*multable)[sti][eni] == -1 {
		(*used)[sti] = true
		for i := 0; i < len((*multable)[sti]); i++ {
			// 不能为-1,1,并且没有使用过
			if (*multable)[sti][i] != -1 && (*multable)[sti][i] != 1 && !isUsed(i, *used) {
				tf, tr := deepFirstSearch(multable, i, eni, used)
				if tf {
					flag = true
					result = (*multable)[sti][i] * tr
				}
			}
		}
		(*used)[sti] = false
	}
	//(*multable)[sti][eni] = result
	//(*multable)[eni][sti] = 1.0 / result
	// 找不到路径的话则返回-1
	return flag, result
}
func isUsed(item int, save map[int]bool) bool {
	result := false
	if key, ok := save[item]; ok && key {
		result = true
	}
	return result
}
