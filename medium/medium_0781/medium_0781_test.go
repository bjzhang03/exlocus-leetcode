package medium_0781

import (
	"fmt"
	"testing"
)

func TestNumRabbits(t *testing.T) {
	var cases = []struct {
		answers  []int
		expected int
	}{
		{[]int{1, 1, 2}, 5},
		{[]int{10, 10, 10}, 11},
		{[]int{}, 0},
		{[]int{0, 0, 1, 1, 1}, 6},
		{[]int{0, 1, 0, 2, 0, 1, 0, 2, 1, 1}, 11},
	}

	for _, val := range cases {
		actual := numRabbits(val.answers)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s\n", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
