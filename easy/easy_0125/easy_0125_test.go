package easy_0125

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	var ip = []struct {
		s        string
		expected bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"race a car", false},
	}

	for _, val := range ip {
		actual := isPalindrome(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
