package medium_0462

import "testing"

func TestMinMoves2(t *testing.T) {

	var mm = []struct {
		nums     []int
		expected int
	}{
		{[]int{1, 2, 4}, 3},
		{[]int{1, 2, 3}, 2},
	}

	for _, val := range mm {
		actual := minMoves2(val.nums)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
