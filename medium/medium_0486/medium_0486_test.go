package medium_0486

import "testing"

func TestPredictTheWinner(t *testing.T) {

	var pt = []struct {
		nums     []int
		expected bool
	}{
		{[]int{1, 5, 2}, false},
		{[]int{1, 5, 233, 7}, true},
	}

	for _, val := range pt {

		actual := PredictTheWinner(val.nums)
		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
