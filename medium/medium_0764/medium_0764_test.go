package medium_0764

import (
	"fmt"
	"testing"
)

func TestOrderOfLargestPlusSign(t *testing.T) {

	var cases = []struct {
		N        int
		mines    [][]int
		expected int
	}{
		{5, [][]int{{4, 2}}, 2},
		{2, [][]int{}, 1},
		{1, [][]int{{0, 0}}, 0},
	}

	for _, val := range cases {
		actual := orderOfLargestPlusSign(val.N, val.mines)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
