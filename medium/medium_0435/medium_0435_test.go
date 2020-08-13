package medium_0435

import "testing"

func TestEraseOverlapIntervals(t *testing.T) {

	var eoi = []struct {
		intervals [][]int
		expected  int
	}{
		{[][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}, 1},
		{[][]int{{1, 2}, {1, 2}, {1, 2}}, 2},
		{[][]int{{1, 2}, {2, 3}}, 0},
	}

	for _, val := range eoi {
		actual := eraseOverlapIntervals(val.intervals)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %d ,expected := %d \n", actual, val.expected)
		}
	}

}
