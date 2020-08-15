package easy_0013

import (
	"fmt"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	var rti = []struct {
		s        string
		expected int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, val := range rti {
		actual := romanToInt(val.s)
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
