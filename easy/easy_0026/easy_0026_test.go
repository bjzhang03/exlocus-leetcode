package easy_0026

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	var rd = []struct {
		nums     []int
		expected int
	}{
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{1, 1, 2}, 2},
	}

	for _, val := range rd {
		actual := removeDuplicates(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
