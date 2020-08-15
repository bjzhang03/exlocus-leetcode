package easy_0058

import (
	"fmt"
	"testing"
)

func TestLengthOfLastWord(t *testing.T) {
	var lolw = []struct {
		s        string
		expected int
	}{
		{"Hello World", 5},
	}

	for _, val := range lolw {
		actual := lengthOfLastWord(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actusl := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
