package medium_0454

import "testing"

func TestFourSumCount(t *testing.T) {

	var fsc = []struct {
		A        []int
		B        []int
		C        []int
		D        []int
		expected int
	}{
		{[]int{0, 1, -1}, []int{-1, 1, 0}, []int{0, 0, 1}, []int{-1, 1, 1}, 17},
		//{[]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2}, 2},
	}

	for _, val := range fsc {
		actual := fourSumCount(val.A, val.B, val.C, val.D)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}
}
