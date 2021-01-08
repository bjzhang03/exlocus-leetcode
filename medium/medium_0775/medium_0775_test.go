package medium_0775

import (
	"fmt"
	"testing"
)

func TestIsIdealPermutation(t *testing.T) {
	var cases = []struct {
		A        []int
		expected bool
	}{
		{[]int{1, 0, 2}, true},
		{[]int{1, 2, 0}, false},
	}
	
	for _, val := range cases {
		actual := isIdealPermutation(val.A)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
