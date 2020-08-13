package medium_0518

import "testing"

func TestChange(t *testing.T) {

	var ch = []struct {
		amount   int
		coins    []int
		expected int
	}{
		{5, []int{1, 2, 5}, 4},
		{0, []int{}, 1},
		{500, []int{1, 2, 5}, 12701},
	}

	for _, val := range ch {
		actual := change(val.amount, val.coins)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}
}
