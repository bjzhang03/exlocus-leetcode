package medium_0075

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {

	var cases = []struct {
		nums     []int // input
		expected []int // expected result
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
	}

	for _, val := range cases {
		sortColors(val.nums)
		if !reflect.DeepEqual(val.nums, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.nums))
		}
	}

}
