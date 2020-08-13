package medium_0718

import "testing"

func TestFindLength(t *testing.T) {
	var fl = []struct {
		A        []int
		B        []int
		expected int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 1, 1, 1, 1}, []int{1, 0, 1, 0, 1}, 2},
	}

	for _, val := range fl {
		actual := findLength(val.A, val.B)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %d, expected := %d", actual, val.expected)
		}
	}

}
