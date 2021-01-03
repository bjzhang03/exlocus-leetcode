package easy_0350

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestIntersect(t *testing.T) {
	var cases = []struct {
		nums1    []int
		nums2    []int
		expected []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}

	for _, val := range cases {
		actual := intersect(val.nums1, val.nums2)
		// 计算两个几个的交集,可以排序再比较
		sort.Ints(actual)
		if sort.Ints(val.expected); !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
