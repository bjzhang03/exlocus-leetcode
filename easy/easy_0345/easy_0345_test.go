package easy_0345

import (
	"fmt"
	"testing"
)

func TestReverseVowels(t *testing.T) {
	var cases = []struct {
		s        string
		expected string
	}{
		{"hello", "holle"},
		{"leetcode", "leotcede"},
	}

	for _, val := range cases {
		actual := reverseVowels(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
