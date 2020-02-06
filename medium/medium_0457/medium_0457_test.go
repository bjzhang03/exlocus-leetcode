package medium_0457

import "testing"

func TestCircularArrayLoop(t *testing.T) {

	var cal = []struct {
		nums     []int
		expected bool
	}{
		{[]int{2, -1, 1, 2, 2}, true},
		{[]int{-1, 2}, false},
		{[]int{-2, 1, -1, -2, -2}, false},
		{[]int{2, 2, 2, 2, 2, 4, 7}, false},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{-1, -1, -1}, true},
		{[]int{-2, -3, -9}, false},
	}

	for _, val := range cal {
		actual := circularArrayLoop(val.nums)

		if actual != val.expected {
			t.Error("Test Failed!", val.nums)
		}
	}

}
