package medium_0377

import "testing"

func TestCombinationSum4(t *testing.T) {

	var cs4 = []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3}, 4, 7},
		{[]int{9}, 3, 0},
		{[]int{1, 2, 4}, 32, 39882198},
		{[]int{}, 1, 0},
	}

	for _, val := range cs4 {
		actual := combinationSum4(val.nums, val.target)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
