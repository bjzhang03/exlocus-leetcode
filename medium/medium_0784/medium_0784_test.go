package medium_0784

import (
	"fmt"
	"sort"
	"testing"
)

func TestLetterCasePermutation(t *testing.T) {
	var cases = []struct {
		S        string
		expected []string
	}{
		{"a1b2", []string{"a1b2", "a1B2", "A1b2", "A1B2"}},
		{"3z4", []string{"3z4", "3Z4"}},
		{"12345", []string{"12345"}},
		{"0", []string{"0"}},
	}

	for _, val := range cases {
		actual := letterCasePermutation(val.S)

		if !strsEqual(val.expected, actual) {
			t.Errorf("Test failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))

		}
	}

}

/*判断两个string的数组是否相等*/
func strsEqual(a, b []string) bool {
	if len(a) == len(b) {
		sort.Strings(a)
		sort.Strings(b)

		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	return false
}
