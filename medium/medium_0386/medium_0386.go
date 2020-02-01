package medium_0386

import (
	"sort"
	"strconv"
)

// 这里调用了默认的库函数进行了处理,不是最优的处理方案
func lexicalOrder(n int) []int {
	save := lexical{}
	for i := 1; i <= n; i++ {
		save.data = append(save.data, i)
	}
	sort.Sort(save)
	return save.data
}

type lexical struct {
	data []int
}

// 获取此slice的长度
func (l lexical) Len() int { return len(l.data) }

// 自定义的排序算法
func (l lexical) Less(i, j int) bool {
	return strconv.Itoa(l.data[i]) < strconv.Itoa(l.data[j])
}

// 交换数据
func (l lexical) Swap(i, j int) {
	tmp := l.data[i]
	l.data[i] = l.data[j]
	l.data[j] = tmp
}
