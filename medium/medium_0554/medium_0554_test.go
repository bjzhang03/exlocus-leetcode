package medium_0554

import "testing"

func TestLeastBricks(t *testing.T) {

	var lb = []struct {
		wall     [][]int
		expected int
	}{
		{[][]int{[]int{1, 2, 2, 1}, []int{3, 1, 2}, []int{1, 3, 2}, []int{2, 4}, []int{3, 1, 2}, []int{1, 3, 1, 1}}, 2},
	}

	for _, val := range lb {

		actual := leastBricks(val.wall)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
