package easy_0172

import (
	"fmt"
	"testing"
)

func TestTrailingZeroes(t *testing.T) {
	var cases = []struct {
		n        int
		expected int
	}{
		{3, 0},
		{5, 1},
	}

	for _, val := range cases {
		actual := trailingZeroes(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))

		}
	}

}
