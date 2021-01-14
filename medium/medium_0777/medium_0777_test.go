package medium_0777

import (
	"fmt"
	"testing"
)

func TestCanTransform(t *testing.T) {
	var cases = []struct {
		start    string
		end      string
		expected bool
	}{
		{"RXXLRXRXL", "XRLXXRRLX", true},
		{"X", "L", false},
		{"LLR", "RRL", false},
		{"XL", "LX", true},
		{"XLLR", "LXLX", false},
		{"XXXXXLXXXX", "LXXXXXXXXX", true},
		{"XXRXLXRXXX", "XXRLXXXXXR", true},
		{"RLX", "XLR", false},
		//{"XLXRRXXRXX", "LXXXXXXRRR", true},
	}

	for _, val := range cases {
		actual := canTransform(val.start, val.end)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
