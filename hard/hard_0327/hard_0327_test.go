package hard_0327

import "testing"

type rsum struct {
	nums  []int
	lower int
	upper int
}

func TestCountRangeSum(t *testing.T) {

	var countersum = []struct {
		in       rsum
		expected int
	}{
		{rsum{[]int{-2, 5, -1}, -2, 2}, 3},
	}

	for _, crs := range countersum {
		actual := countRangeSum(crs.in.nums, crs.in.lower, crs.in.upper)

		if actual != crs.expected {
			t.Errorf("test failed!")

		}
	}

}
