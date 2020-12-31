package easy_0258

import (
	"fmt"
	"testing"
)

func TestAddDigits(t *testing.T) {

	var cases = []struct {
		num      int
		expected int
	}{
		{38, 2},
	}

	for _, val := range cases {
		actual := addDigits(val.num)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
