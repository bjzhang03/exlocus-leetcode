package medium_0343

import "testing"

func TestIntegerBreak(t *testing.T) {

	var integerBreakTest = []struct {
		in       int
		expected int
	}{
		{2, 1},
		{3, 2},
		{10, 36},
		{8, 18},
	}

	for _, val := range integerBreakTest {
		if integerBreak(val.in) != val.expected {
			t.Errorf("Test Error!")
		}
	}

}
