package medium_0576

import (
	"testing"
)

func TestFindPaths(t *testing.T) {

	var fp = []struct {
		m        int
		n        int
		N        int
		i        int
		j        int
		expected int
	}{
		{2, 2, 2, 0, 0, 6},
		{1, 3, 3, 0, 1, 12},
		{7, 6, 13, 0, 2, 1659429},
		{8, 50, 23, 5, 26, 914783380},
		{36, 5, 50, 15, 3, 390153306},
	}

	for _, val := range fp {
		actual := findPaths(val.m, val.n, val.N, val.i, val.j)

		if actual != val.expected {
			t.Errorf("Test Failed! actual:= %d, expected:= %d", actual, val.expected)
		}
	}

}

/***
3
10
34
119
418
1490
5324
19184
69216
250953
910284
3311387
12045841
--- PASS: TestFindPaths (5.49s)
PASS
 */
