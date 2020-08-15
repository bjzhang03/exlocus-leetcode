package easy_0088

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var me = []struct {
		nums1    []int
		m        int
		nums2    []int
		n        int
		expected []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
	}
	
	for _, val := range me {
		merge(val.nums1, val.m, val.nums2, val.n)

		if !reflect.DeepEqual(val.nums1, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.nums1))
		}
	}

}
