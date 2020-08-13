package easy_0001

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {

	var ts = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, val := range ts {
		actual := twoSum(val.nums, val.target)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed!")
		}
	}

}
