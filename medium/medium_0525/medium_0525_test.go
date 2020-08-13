package medium_0525

import "testing"

func TestFindMaxLength(t *testing.T) {

	var fml = []struct {
		nums     []int
		expected int
	}{
		{[]int{0, 1}, 2},
		{[]int{0, 1, 0}, 2},
		{[]int{0, 0, 1, 0, 0, 1, 0}, 4},
		{[]int{0, 1, 1}, 2},
	}

	for _, val := range fml {
		actual := findMaxLength(val.nums)
		if actual != val.expected {
			t.Error("Test Failed!", val.nums, val.expected, actual)
		}
	}

}
