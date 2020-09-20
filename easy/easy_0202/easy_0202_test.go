package easy_0202

import (
	"fmt"
	"testing"
)

func TestIsHappy(t *testing.T) {

	var cases = []struct {
		n        int
		expected bool
	}{
		{19, true},
	}

	for _, val := range cases {
		actual := isHappy(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
