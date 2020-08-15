package easy_0009

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var ip = []struct {
		x        int
		expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
	}

	for _, val := range ip {
		actual := isPalindrome(val.x)
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
