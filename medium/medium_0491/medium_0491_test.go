package medium_0491

import (
	"reflect"
	"testing"
)

func TestFindSubsequences(t *testing.T) {

	var fs = []struct {
		nums     []int
		expected [][]int
	}{
		//{[]int{4, 6, 7, 7}, [][]int{{4, 6}, {4, 7}, {4, 6, 7}, {4, 6, 7, 7}, {6, 7}, {6, 7, 7}, {7, 7}, {4, 7, 7}}},
	}

	for _, val := range fs {
		actual := findSubsequences(val.nums)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Error("Test Failed!")
		}
	}

}
