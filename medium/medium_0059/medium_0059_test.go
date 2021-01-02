package medium_0059

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateMatrix(t *testing.T) {
	var cases = []struct {
		n        int
		expected [][]int
	}{
		{3, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{1, [][]int{{1}}},
	}

	for _, val := range cases {
		actual := generateMatrix(val.n)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
