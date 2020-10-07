package easy_0205

import (
	"fmt"
	"testing"
)

func TestIsIsomorphic(t *testing.T) {
	var cases = []struct {
		s        string
		t        string
		expected bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}

	for _, val := range cases {
		actual := isIsomorphic(val.s, val.t)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
