package easy_0189

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	var cases = []struct {
		nums     []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for _, val := range cases {
		rotate(val.nums, val.k)

		if !reflect.DeepEqual(val.expected, val.nums) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.nums))
		}

	}
}
