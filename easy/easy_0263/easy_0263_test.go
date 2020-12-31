package easy_0263

import (
	"fmt"
	"testing"
)

func TestIsUgly(t *testing.T) {

	var cases = []struct {
		num      int
		expected bool
	}{
		{6, true},
		{8, true},
		{14, false},
	}

	for _, val := range cases {
		actual := isUgly(val.num)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
