package medium_0565

import "testing"

func TestArrayNesting(t *testing.T) {
	var an = []struct {
		nums     []int
		expected int
	}{
		{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
	}

	for _, val := range an {
		actual := arrayNesting(val.nums)

		if actual != val.expected {
			t.Errorf("Test Failed! actual:= %d, expected:= %d", actual, val.expected)
		}
	}

}
