package medium_0494

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	var ftsw = []struct {
		nums     []int
		S        int
		expected int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 1}, 1, 256},
		{[]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}, 0, 1048576},
	}

	for _, val := range ftsw {

		actual := findTargetSumWays(val.nums, val.S)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
