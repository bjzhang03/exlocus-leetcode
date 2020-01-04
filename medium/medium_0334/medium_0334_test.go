package medium_0334

import "testing"

func TestIncreasingTriplet(t *testing.T) {

	var icr = []struct {
		in       []int
		expected bool
	}{
		{[]int{5, 1, 5, 5, 2, 5, 4}, true},
	}

	for _, val := range icr {
		actual := increasingTriplet(val.in)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
