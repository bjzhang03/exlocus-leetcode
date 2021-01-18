package medium_0785

import (
	"fmt"
	"testing"
)

func TestIsBipartite(t *testing.T) {
	var cases = []struct {
		graph    [][]int
		expected bool
	}{
		{[][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}, true},
		{[][]int{{1, 2, 3}, {0, 2}, {0, 1, 3}, {0, 2}}, false},
		{[][]int{{4, 1}, {0, 2}, {1, 3}, {2, 4}, {3, 0}}, false},
		{[][]int{{3}, {2, 4}, {1}, {0, 4}, {1, 3}}, true},
		{[][]int{{}, {2, 4, 6}, {1, 4, 8, 9}, {7, 8}, {1, 2, 8, 9}, {6, 9}, {1, 5, 7, 8, 9}, {3, 6, 9}, {2, 3, 4, 6, 9}, {2, 4, 5, 6, 7, 8}}, false},
		{[][]int{{}, {3}, {}, {1}, {}}, true},
	}

	for _, val := range cases {
		actual := isBipartite(val.graph)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s\n", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
