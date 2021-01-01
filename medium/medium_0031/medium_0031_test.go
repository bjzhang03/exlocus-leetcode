package medium_0031

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected []int // expected result
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1}, []int{1}},
	}

	for _, val := range cases {
		nextPermutation(val.nums)
		if !reflect.DeepEqual(val.nums, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.nums))
		}
	}
}
