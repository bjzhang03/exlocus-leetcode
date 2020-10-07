package easy_0231

import (
	"fmt"
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {

	var cases = []struct {
		n        int
		expected bool
	}{
		{-1, false},
		{0, false},
		{1, true},
		{16, true},
		{3, false},
		{4, true},
		{5, false},
	}

	for _, val := range cases {
		actual := isPowerOfTwo(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
