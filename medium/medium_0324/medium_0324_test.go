package medium_0324

import (
	"reflect"
	"testing"
)

func TestWiggleSort(t *testing.T) {
	var wiggleTest = []struct {
		in       []int
		expected []int
	}{
		{[]int{9, 11, 3, 5, 1, 5, 6, 8, 9, 11}, []int{1, 5, 1, 1, 6, 4}},
		{[]int{1, 1, 2, 1, 2, 2, 1}, []int{1, 5, 1, 1, 6, 4}},
		{[]int{1, 5, 1, 1, 6, 4}, []int{1, 5, 1, 1, 6, 4}},
		{[]int{4, 5, 5, 6}, []int{1, 5, 1, 1, 6, 4}},
	}

	for _, val := range wiggleTest {
		wiggleSort(val.in)

		if !reflect.DeepEqual(val.in, val.expected) {

		}
	}
}
