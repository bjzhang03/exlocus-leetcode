package easy_0038

import (
	"fmt"
	"testing"
)

func TestCountAndSay(t *testing.T) {
	var cas = []struct {
		n        int
		expected string
	}{
		{1, "1"},
		{4, "1211"},
	}

	for _, val := range cas {
		actual := countAndSay(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actusl := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
