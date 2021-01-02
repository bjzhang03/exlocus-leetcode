package medium_0056

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	var cases = []struct {
		intervals [][]int
		expected  [][]int
	}{
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
		{[][]int{{1, 4}, {1, 4}}, [][]int{{1, 4}}},
		{[][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}}, [][]int{{1, 3}, {4, 7}}},
	}

	for _, val := range cases {
		actual := merge(val.intervals)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
