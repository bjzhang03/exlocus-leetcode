package medium_0769

import (
	"fmt"
	"testing"
)

func TestMaxChunksToSorted(t *testing.T) {

	var cases = []struct {
		arr      []int
		expected int
	}{
		{[]int{4, 3, 2, 1, 0}, 1},
		{[]int{1, 0, 2, 3, 4}, 4},
	}

	for _, val := range cases {
		actual := maxChunksToSorted(val.arr)

		if actual != val.expected {
			t.Errorf("Test failed! expected := %s, actual := %s\n", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
