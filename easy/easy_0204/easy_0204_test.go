package easy_0204

import (
	"fmt"
	"testing"
)

func TestCountPrimes(t *testing.T) {
	var cases = []struct {
		n        int
		expected int
	}{
		{10, 4},
	}

	for _, val := range cases {
		actual := countPrimes(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
