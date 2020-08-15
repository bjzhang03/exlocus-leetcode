package easy_0053

import (
	"fmt"
	"testing"
)

func TestMaxSubArray(t *testing.T) {
	var msa = []struct {
		nums     []int
		expected int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for _, val := range msa {
		actual := maxSubArray(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
