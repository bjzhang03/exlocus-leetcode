package medium_0148

import "testing"

func TestSortList(t *testing.T) {
	head := buildListFromSlice([]int{4, 2, 1, 3})

	sortList(head)
}
