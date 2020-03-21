package medium_0592

import "testing"

func TestFractionAddition(t *testing.T) {
	var fa = []struct {
		expression string

		expected string
	}{
		{"-1/2+1/2", "0/1"},
		{"-1/2+1/2+1/3", "1/3"},
		{"1/3-1/2", "-1/6"},
		{"5/3+1/3", "2/1"},
		{"7/3+5/2-3/10", "68/15"},
		{"8/3+1/1-8/3", "1/1"},
	}

	for _, val := range fa {
		actual := fractionAddition(val.expression)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %s, expected := %s \n", actual, val.expected)
		}
	}

}
