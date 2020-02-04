package medium_0433

var genes = []string{"A", "C", "G", "T"}

func minMutation(start string, end string, bank []string) int {
	// 不可达的话,默认为-1
	result := -1
	if len(bank) > 0 {
		result = 0
		// 预先分配好空间大小
		// 将bank数据转化成map
		bankmap := make(map[string]bool, len(bank)+2)
		for i := 0; i < len(bank); i++ {
			bankmap[bank[i]] = true
		}

		// 如果map无法到达end,则直接返回-1
		if _, ok := bankmap[end]; !ok {
			return -1
		} else {
			// 广度优先搜索
			// 记录下使用过的记录
			used := make(map[string]bool, len(bank)+2)
			cset := map[string]bool{start: true}
			for {
				// 找到下一次可达的数据
				nset := map[string]bool{}
				for key, _ := range cset {
					sliceNcs := getGeneFromBank(key, bankmap, used)
					for i := 0; i < len(sliceNcs); i++ {
						nset[sliceNcs[i]] = true
					}
					used[key] = true
				}
				if result > 10 {
					break
				}
				// 更新一下结果
				if _, ousedk := used[end]; ousedk {
					return result
				}
				// 开始下一次得计算处理
				result++
				// 判断下一次的数据是否可达
				if len(nset) == 0 {
					return -1
				} else {
					cset = nset
				}

			}
		}
	}
	return result
}

func getGeneFromBank(gene string, bank, used map[string]bool) []string {
	result := []string{}
	for i := 0; i < len(gene); i++ {
		cgene := gene[:i]
		for j := 0; j < len(genes); j++ {
			if _, ou := used[cgene+genes[j]+gene[i+1:]]; !ou {
				if _, ob := bank[cgene+genes[j]+gene[i+1:]]; ob {
					result = append(result, cgene+genes[j]+gene[i+1:])
				}
			}
		}
	}
	return result
}
