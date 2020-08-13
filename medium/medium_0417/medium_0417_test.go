package medium_0417

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {

	var pa = []struct {
		matrix   [][]int
		expected [][]int
	}{
		{[][]int{{1, 2, 2, 3, 5}, {3, 2, 3, 4, 4}, {2, 4, 5, 3, 1}, {6, 7, 1, 4, 5}, {5, 1, 1, 2, 4}}, [][]int{{0, 4}, {1, 3}, {1, 4}, {2, 2}, {3, 0}, {3, 1}, {4, 0}}},
		//{[][]int{{1, 3, 5}, {3, 4, 4}}, [][]int{}},
	}

	for _, val := range pa {
		actual := pacificAtlantic(val.matrix)
		if !reflect.DeepEqual(actual, val.expected) {
			t.Error("Test Failed!")
		}
	}
}
