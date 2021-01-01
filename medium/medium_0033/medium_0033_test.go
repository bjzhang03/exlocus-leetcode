package medium_0033

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	var cases = []struct {
		nums     []int
		target   int
		expected int // expected result
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
	}

	for _, val := range cases {
		actual := search(val.nums, val.target)
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
