package easy_0136

import (
	"fmt"
	"testing"
)

func TestSingleNumber(t *testing.T) {
	var sn = []struct {
		nums     []int
		expected int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
	}

	for _, val := range sn {
		actual := singleNumber(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
