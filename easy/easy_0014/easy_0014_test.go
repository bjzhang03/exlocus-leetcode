package easy_0014

import (
	"fmt"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	var lcp = []struct {
		strs     []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for _, val := range lcp {
		actual := longestCommonPrefix(val.strs)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
