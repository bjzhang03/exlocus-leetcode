package medium_0738

import "testing"

func TestMonotoneIncreasingDigits(t *testing.T) {

	var mid = []struct {
		N        int
		expected int
	}{
		{10, 9},
		{1234, 1234},
		{332, 299},
	}

	for _, val := range mid {
		actual := monotoneIncreasingDigits(val.N)

		if actual != val.expected {
			t.Errorf("Test Failed! input := %d, expected := %d, actual := %d", val.N, val.expected, actual)
		}
	}

}
