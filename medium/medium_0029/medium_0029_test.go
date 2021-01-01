package medium_0029

import (
	"fmt"
	"testing"
)

func TestDivide(t *testing.T) {
	var cases = []struct {
		dividend int
		divisor  int
		expected int // expected result
	}{
		{10, 3, 3},
		{7, -3, -2},
		{-2147483648, -1, 2147483647},
	}

	for _, val := range cases {
		actual := divide(val.dividend, val.divisor)
		if val.expected != actual {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
