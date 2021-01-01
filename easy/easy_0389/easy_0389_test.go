package easy_0389

import (
	"fmt"
	"testing"
)

func TestFindTheDifference(t *testing.T) {
	var cases = []struct {
		s        string
		t        string
		expected byte
	}{
		{"abcd", "abcde", 'e'},
		{"", "y", 'y'},
		{"a", "aa", 'a'},
		{"ae", "aea", 'a'},
	}

	for _, val := range cases {
		actual := findTheDifference(val.s, val.t)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
