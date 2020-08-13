package medium_0678

import "testing"

func TestCheckValidString(t *testing.T) {
	var cvs = []struct {
		s        string
		expected bool
	}{
		{"()", true},
		{"(*)", true},
		{"(*))", true},
		{"*", true},
	}

	for _, val := range cvs {
		actual := checkValidString(val.s)
		if actual != val.expected {
			t.Errorf("Test Failed! s := %s, actual:=%t, expected:=%t", val.s, actual, val.expected)
		}
	}

}
