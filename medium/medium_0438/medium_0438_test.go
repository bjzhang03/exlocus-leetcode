package medium_0438

import (
	"reflect"
	"testing"
)

func TestFindAnagrams(t *testing.T) {

	var fas = []struct {
		s        string
		p        string
		expected []int
	}{
		{"cbaebabacd", "abc", []int{0, 6}},
		{"abab", "ab", []int{0, 1, 2}},
	}

	for _, val := range fas {
		actual := findAnagrams(val.s, val.p)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Error("Test Failed!s")
		}
	}

}
