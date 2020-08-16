package easy_0171

import (
	"fmt"
	"testing"
)

func TestTitleToNumber(t *testing.T) {

	var cases = []struct {
		s        string
		expected int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for _, val := range cases {
		actual := titleToNumber(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
