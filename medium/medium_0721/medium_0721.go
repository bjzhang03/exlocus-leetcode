package medium_0721

import (
	"fmt"
	"sort"
	"strings"
)

func accountsMerge(accounts [][]string) [][]string {
	if len(accounts) > 0 {
		// 先使用广度优先的思想进行深度优先遍历获得所有需要合并到一起的数据
		current := [][]string{accounts[0]}
		used := map[string]bool{fmt.Sprint(accounts[0]): true}
		currentres := [][]string{}
		for len(current) > 0 {
			currentres = append(currentres, current...)
			ncurrent := [][]string{}
			for i := 0; i < len(current); i++ {
				for j := 0; j < len(current[i]); j++ {
					if strings.Contains(current[i][j], "@") {
						currentmail := current[i][j]
						// 在数据中找出含有当前mail的数据,然后添加进来
						// 同时将数据标记为已经使用过了
						for k := 0; k < len(accounts); k++ {
							for m := 0; m < len(accounts[k]); m++ {
								if currentmail == accounts[k][m] {
									if _, ok := used[fmt.Sprint(accounts[k])]; !ok {
										ncurrent = append(ncurrent, accounts[k])
										used[fmt.Sprint(accounts[k])] = true
										break
									}
								}
							}
						}
					}
				}
			}
			current = current[:0]
			current = append(current, ncurrent...)
		}
		// 将currentres里面的数据进行合并处理
		usedres := map[string]bool{}
		unsortstr := []string{}
		// 获取数据中的name
		name := ""
		for i := 0; i < len(currentres); i++ {
			for j := 0; j < len(currentres[i]); j++ {
				if strings.Contains(currentres[i][j], "@") {
					if _, ok := usedres[currentres[i][j]]; !ok {
						unsortstr = append(unsortstr, currentres[i][j])
						usedres[currentres[i][j]] = true
					}
				} else if name == "" {
					// name一般都是一样的
					name = currentres[i][j]
				}
			}
		}
		// 对合并后的数据进行排序操作
		sort.Strings(unsortstr)
		result := append([][]string{}, append([]string{name}, unsortstr...))
		// 构造新的需要递归操作的数据
		naccounts := [][]string{}
		for i := 0; i < len(accounts); i++ {
			if _, ok := used[fmt.Sprint(accounts[i])]; !ok {
				naccounts = append(naccounts, accounts[i])
			}
		}
		// 这个地方的递归可以使用for循环替换掉的方式进行优化操作
		return append(result, accountsMerge(naccounts)...)
	}
	return nil
}
