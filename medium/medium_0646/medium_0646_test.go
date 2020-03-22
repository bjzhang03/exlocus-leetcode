package medium_0646

import "testing"

func TestFindLongestChain(t *testing.T) {
	var flc = []struct {
		pairs    [][]int
		expected int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}}, 2},
	}

	for _, val := range flc {
		actual := findLongestChain(val.pairs)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
