package easy_0198

import (
	"fmt"
	"testing"
)

func TestRob(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}

	for _, val := range cases {
		actual := rob(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
