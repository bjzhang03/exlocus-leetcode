package medium_0495

import "testing"

func TestFindPoisonedDuration(t *testing.T) {
	var fpd = []struct {
		timeSeries []int
		duration   int
		expected   int
	}{
		{[]int{1, 4}, 2, 4},
		{[]int{1, 2}, 2, 3},
	}

	for _, val := range fpd {
		actual := findPoisonedDuration(val.timeSeries, val.duration)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}
}
