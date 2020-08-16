package easy_0168

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(t *testing.T) {

	var cases = []struct {
		n        int
		expected string
	}{
		{1, "A"},
		{28, "AB"},
		{701, "ZY"},
	}

	for _, val := range cases {
		actual := convertToTitle(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
