package hard_0126

func sliceHas(used []string, item string) bool {
	for _, value := range used {
		if value == item {
			return true
		}
	}
	return false
}

// bfs查找所有的数据
func buildLadderTree(beginWord string, endWord string, wordList map[string]bool, help *[]map[uint8]bool, used []string, result *[][]string) {
	// 新建树的root
	// 如果已经到达则直接返回即可
	if beginWord == endWord {
		used = append(used, beginWord)
		// 存在数据则进行比较操作
		if len(*result) > 0 {
			tmp := (*result)[0]
			if len(tmp) > len(used) {
				// 清空数据
				(*result) = (*result)[:0]
				tmpr := []string{}
				tmpr = append(tmpr, used...)
				(*result) = append(*result, tmpr)
			} else if len(tmp) == len(used) {
				tmpr := []string{}
				tmpr = append(tmpr, used...)
				(*result) = append(*result, tmpr)
			}
		} else {
			tmpr := []string{}
			tmpr = append(tmpr, used...)
			(*result) = append(*result, tmpr)
		}
	}
	used = append(used, beginWord)
	// 找出可以替换的数据,然后进行替换
	for i := 0; i < len(beginWord); i++ {
		// 找出替换的元素来
		for key, _ := range (*help)[i] {
			// 替换操作
			if key != beginWord[i] {
				nextval := beginWord[:i]
				nextval = nextval + string(key)
				nextval = nextval + beginWord[i+1:]
				// 没有使用过,并且存在wordlist中
				ok1 := sliceHas(used, nextval);
				ok2 := wordList[nextval];
				if !ok1 && ok2 {
					buildLadderTree(nextval, endWord, wordList, help, used, result)
				}
			}
		}
	}
}

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	// 先对数据进行一些预处理
	cletters := []map[uint8]bool{}
	// 找出各个位置使用过的letter数据
	for i := 0; i < len(beginWord); i++ {
		usedLetter := map[uint8]bool{}
		for _, val := range wordList {
			usedLetter[val[i]] = true
		}
		cletters = append(cletters, usedLetter)
	}
	wordListSet := map[string]bool{}
	for _, val := range wordList {
		wordListSet[val] = true
	}
	result := [][]string{}
	buildLadderTree(beginWord, endWord, wordListSet, &cletters, []string{}, &result)
	return result
}
