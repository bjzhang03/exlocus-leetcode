package medium_0127

// bfs查找所有的数据
func breadthFirstSearch(beginWord string, endWord string, wordList map[string]bool, help *[]map[uint8]bool) int {
	// 设置好大小
	queue := make([]string, len(wordList))
	// 数据初始化大小先处理好
	nextQueue := make([]string, len(wordList))
	queue = append(queue, beginWord)
	// 记录相关数据
	result := 1
	// 初始化的时候指定大小
	usedWords := make(map[string]bool, len(wordList))
	usedWords[beginWord] = true
	for len(queue) > 0 {
		// fmt.Println(queue, beginWord)
		nextQueue = nextQueue[:0]
		// 对queue中的数据进行遍历
		for _, val := range queue {
			// 看val中的每一个元素是否都是可以替换的
			for i := 0; i < len(val); i++ {
				// 找出替换的元素来
				for key, _ := range (*help)[i] {
					// 替换操作
					if key != val[i] {
						nextval := val[:i]
						nextval = nextval + string(key)
						nextval = nextval + val[i+1:]
						// 找到则直接返回即可
						if nextval == endWord {
							result++
							return result
						}
						// 没有使用过,并且存在wordlist中
						ok1 := usedWords[nextval]
						ok2 := wordList[nextval]
						if !ok1 && ok2 {
							nextQueue = append(nextQueue, nextval)
							usedWords[nextval] = true
						}
					}
				}
			}
		}
		result++
		// 原地清空数组大小
		queue = queue[:0]
		queue = append(queue, nextQueue...)
	}
	return 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
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
	return breadthFirstSearch(beginWord, endWord, wordListSet, &cletters)
}
