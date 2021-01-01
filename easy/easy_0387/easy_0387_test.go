package easy_0387

import (
	"fmt"
	"testing"
)

func TestFirstUniqChar(t *testing.T) {
	var cases = []struct {
		s        string
		expected int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
	}

	for _, val := range cases {
		actual := firstUniqChar(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
