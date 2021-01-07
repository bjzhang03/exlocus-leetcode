package medium_0767

import (
	"fmt"
	"testing"
)

func TestReorganizeString(t *testing.T) {

	var cases = []struct {
		S string
	}{
		{"aab"},
		{"aaab"},
		{"abbabbaaab"},
		{"bfrbs"},
	}

	for _, val := range cases {
		actual := reorganizeString(val.S)
		//fmt.Println(actual)
		if !check(actual) {
			t.Errorf("Test Failed! actual:= %s", fmt.Sprintf(actual))

		}
	}
}
