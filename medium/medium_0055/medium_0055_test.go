package medium_0055

import (
	"fmt"
	"testing"
)

func TestCanJump(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, 3, 1, 1, 4}, true},
		{[]int{3, 2, 1, 0, 4}, false},
	}

	for _, val := range cases {
		actual := canJump(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
