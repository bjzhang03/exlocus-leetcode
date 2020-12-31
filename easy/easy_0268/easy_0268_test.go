package easy_0268

import (
	"fmt"
	"testing"
)

func TestMissingNumber(t *testing.T) {
	var cases = []struct {
		nums     []int
		expected int
	}{
		{[]int{3, 0, 1}, 2},
		{[]int{0, 1}, 2},
		{[]int{9, 6, 4, 2, 3, 5, 7, 0, 1}, 8},
		{[]int{0}, 1},
	}

	for _, val := range cases {
		actual := missingNumber(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
