package easy_0242

import (
	"fmt"
	"testing"
)

func TestIsAnagram(t *testing.T) {

	var cases = []struct {
		s        string
		t        string
		expected bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, val := range cases {
		actual := isAnagram(val.s, val.t)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprintf("%t", val.expected), fmt.Sprintf("%t", actual))
		}
	}

}
