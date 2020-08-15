package easy_0028

import (
	"fmt"
	"testing"
)

func TestStrStr(t *testing.T) {
	var ss = []struct {
		haystack string
		needle   string
		expected int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}

	for _, val := range ss {
		actual := strStr(val.haystack, val.needle)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
