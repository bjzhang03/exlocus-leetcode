package medium_0779

import (
	"fmt"
	"testing"
)

func TestKthGrammar(t *testing.T) {

	var cases = []struct {
		N        int
		K        int
		expected int
	}{
		{1, 1, 0},
		{2, 1, 0},
		{2, 2, 1},
		{4, 5, 1},
		{3, 3, 1},
	}

	for _, val := range cases {
		actual := kthGrammar(val.N, val.K)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s\n", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
