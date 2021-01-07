package medium_0754

import (
	"fmt"
	"testing"
)

func TestReachNumber(t *testing.T) {

	var cases = []struct {
		target   int
		expected int
	}{
		{3, 2},
		{2, 3},
		{100, 15},
		//{-1000000000, 3},
	}

	for _, val := range cases {
		actual := reachNumber(val.target)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
