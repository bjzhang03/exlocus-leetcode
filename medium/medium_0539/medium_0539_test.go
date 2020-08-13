package medium_0539

import "testing"

func TestFindMinDifference(t *testing.T) {

	var fmd = []struct {
		timePoints []string
		expected   int
	}{
		{[]string{"23:59", "00:00"}, 1},
	}

	for _, val := range fmd {
		if val.expected != findMinDifference(val.timePoints) {
			t.Error("Test Failed!")
		}
	}

}
