package easy_0020

import (
	"fmt"
	"testing"
)

func TestIsValid(t *testing.T) {
	var iv = []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"(]", false},
		{"([)]", false},
		{"{[]}", true},
	}

	for _, val := range iv {
		actual := isValid(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
