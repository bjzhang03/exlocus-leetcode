package medium_0002

import "strconv"

//新建一个listnode的类型数据
type ListNode struct {
	Val  int
	Next *ListNode
}

func (p *ListNode) String() string {
	result := ""
	for it := p; it != nil; it = it.Next {
		result = result + strconv.Itoa(it.Val) + " "

	}
	return result
}
