package medium_0057

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsert(t *testing.T) {
	var cases = []struct {
		intervals   [][]int
		newInterval []int
		expected    [][]int
	}{
		{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
		{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
		{[][]int{}, []int{5, 7}, [][]int{{5, 7}}},
		{[][]int{{1, 5}}, []int{2, 3}, [][]int{{1, 5}}},
		{[][]int{{1, 5}}, []int{2, 7}, [][]int{{1, 7}}},
	}

	for _, val := range cases {
		actual := insert(val.intervals, val.newInterval)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
