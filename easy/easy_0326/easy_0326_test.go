package easy_0326

import (
	"fmt"
	"testing"
)

func TestIsPowerOfThree(t *testing.T) {
	var cases = []struct {
		n        int
		expected bool
	}{
		{27, true},
		{0, false},
		{9, true},
		{45, false},
	}

	for _, val := range cases {
		actual := isPowerOfThree(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
