package medium_0695

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {

	var maoi = []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 1, 1}, {0, 0, 0, 1, 1}}, 4},
	}

	for _, val := range maoi {
		actual := maxAreaOfIsland(val.grid)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %d ,expected := %d", actual, val.expected)
		}
	}

}
