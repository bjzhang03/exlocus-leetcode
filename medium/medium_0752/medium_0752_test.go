package medium_0752

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {

	var cases = []struct {
		deadends []string
		target   string
		expected int
	}{
		{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{[]string{"8888"}, "0009", 1},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
		{[]string{"0000"}, "8888", -1},
	}

	for _, val := range cases {
		actual := openLock(val.deadends, val.target)

		if actual != val.expected {
			t.Errorf("Test failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
