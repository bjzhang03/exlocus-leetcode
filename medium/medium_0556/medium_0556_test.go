package medium_0556

import (
	"testing"
)

func TestNextGreaterElement(t *testing.T) {

	var nge = []struct {
		n        int
		expected int
	}{
		{12, 21},
		{21, -1},
		{2, -1},
		{312, 321},
		{1234, 1243},
		{230241, 230412},
		{12443322, 13222344},
		{1999999999, -1},
	}

	for _, val := range nge {
		actual := nextGreaterElement(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! actual:= %d, expected:= %d", actual, val.expected)
		}
	}

}
