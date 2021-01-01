package easy_0383

import (
	"fmt"
	"testing"
)

func TestCanConstruct(t *testing.T) {
	var cases = []struct {
		ransomNote string
		magazine   string
		expected   bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}

	for _, val := range cases {
		actual := canConstruct(val.ransomNote, val.magazine)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
