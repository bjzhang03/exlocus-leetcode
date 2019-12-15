package medium_0310

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

type tree struct {
	n     int
	edges [][]int
}

func TestFindMinHeightTrees(t *testing.T) {

	var findTests = []struct {
		in       tree
		expected []int
	}{
		{tree{4, [][]int{{1, 0}, {1, 2}, {1, 3}}}, []int{1}},
		{tree{6, [][]int{{0, 3}, {1, 3}, {2, 3}, {4, 3}, {5, 4}}}, []int{3, 4}},
		{tree{6, [][]int{{0, 1}, {0, 2}, {0, 3}, {3, 4}, {4, 5}}}, []int{3}},
	}

	for _, ft := range findTests {
		actual := findMinHeightTrees(ft.in.n, ft.in.edges)
		sort.Ints(actual)
		sort.Ints(ft.expected)

		fmt.Println(actual)
		fmt.Println(ft.expected)
		if !reflect.DeepEqual(actual, ft.expected) {
			t.Errorf("Test failed")
		}
	}

}
