package medium_0523

import "testing"

func TestCheckSubarraySum(t *testing.T) {

	var css = []struct {
		nums     []int
		k        int
		expected bool
	}{
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 4, 6, 7}, -6, true},
		{[]int{0, 0}, 0, true},
	}

	for _, val := range css {
		actual := checkSubarraySum(val.nums, val.k)
		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
