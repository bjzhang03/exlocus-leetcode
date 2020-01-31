package medium_0375

import "testing"

func TestGetMoneyAmount(t *testing.T) {

	var gma = []struct {
		in       int
		expected int
	}{
		{1, 0},
		{2, 1},
		{3, 2},
		{4, 4},
		{5, 6},
		{6, 8},
		{7, 10},
		{8, 12},
	}

	for _, val := range gma {
		actual := getMoneyAmount(val.in)

		if actual != val.expected {
			t.Errorf("Test failed, Input: %d, Actual: %d, Expected: %d \n", val.in, actual, val.expected)
		}
	}

}
