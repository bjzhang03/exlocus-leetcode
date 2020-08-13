package medium_0684

import (
	"reflect"
	"testing"
)

func TestFindRedundantConnection(t *testing.T) {

	var frc = []struct {
		edges    [][]int
		expected []int
	}{
		{[][]int{{1, 2}, {1, 3}, {2, 3}}, []int{2, 3}},
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 4}, {1, 5}}, []int{1, 4}},
		{[][]int{{3, 4}, {1, 2}, {2, 4}, {3, 5}, {2, 5}}, []int{2, 5}},
	}

	for _, val := range frc {
		actual := findRedundantConnection(val.edges)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! actual := %d , expected := %d", actual, val.expected)
		}
	}

}
