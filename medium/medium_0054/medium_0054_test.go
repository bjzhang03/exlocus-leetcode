package medium_0054

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	var cases = []struct {
		matrix   [][]int
		expected []int
	}{
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}

	for _, val := range cases {
		actual := spiralOrder(val.matrix)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
