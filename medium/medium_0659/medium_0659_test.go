package medium_0659

import "testing"

func TestIsPossible(t *testing.T) {
	var ip = []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 2, 3, 3, 4, 5}, true},
		{[]int{1, 2, 3, 3, 4, 4, 5, 5}, true},
		{[]int{1, 2, 3, 4, 4, 5}, false},
		{[]int{1, 2, 3, 3, 4, 4, 5, 5}, true},
	}

	for _, val := range ip {
		actual := isPossible(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! nums := %d, actual := %t, expected := %t", val.nums, actual, val.expected)
		}
	}

}
