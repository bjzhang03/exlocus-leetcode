package medium_0503

import (
	"reflect"
	"testing"
)

func TestNextGreaterElements(t *testing.T) {

	var nge = []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 1}, []int{2, -1, 2}},
	}

	for _, val := range nge {
		actual := nextGreaterElements(val.nums)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Error("Test Failed!")
		}
	}

}
