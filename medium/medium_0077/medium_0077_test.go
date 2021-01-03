package medium_0077

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombine(t *testing.T) {
	var cases = []struct {
		n        int
		k        int
		expected [][]int // expected result
	}{
		{4, 2, [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}}},
		{1, 1, [][]int{{1}}},
	}

	for _, val := range cases {
		actual := combine(val.n, val.k)
		if !sliceSetEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}

func sliceSetEqual(a [][]int, b [][]int) bool {
	if len(a) == len(b) {
		setcontainsf := func(set [][]int, item []int) bool {
			for i := 0; i < len(set); i++ {
				if reflect.DeepEqual(set[i], item) {
					return true
				}
			}
			return false
		}
		/*判断所有的a都在b中*/
		for i := 0; i < len(a); i++ {
			if !setcontainsf(b, a[i]) {
				return false
			}
		}
		/*判断所有的b都在a中*/
		for i := 0; i < len(b); i++ {
			if !setcontainsf(a, b[i]) {
				return false
			}
		}
		return true
	}
	return false
}
