package easy_0067

import (
	"fmt"
	"testing"
)

func TestAddBinary(t *testing.T) {
	var ab = []struct {
		a        string
		b        string
		expected string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}

	for _, val := range ab {
		actual := addBinary(val.a, val.b)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
