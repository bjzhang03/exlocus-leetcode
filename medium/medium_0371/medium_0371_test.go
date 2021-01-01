package medium_0371

import (
	"fmt"
	"testing"
)

func TestGetSum(t *testing.T) {
	var cases = []struct {
		a        int
		b        int
		expected int
	}{
		{1, 2, 3},
		{-2, 3, 1},
	}

	for _, val := range cases {
		actual := getSum(val.a, val.b)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
