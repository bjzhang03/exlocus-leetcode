package easy_0283

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
	}

	for _, val := range cases {
		moveZeroes(val.nums)

		if !reflect.DeepEqual(val.nums, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.nums))
		}
	}
}
