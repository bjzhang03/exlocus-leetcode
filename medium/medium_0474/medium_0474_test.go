package medium_0474

import "testing"

func TestFindMaxForm(t *testing.T) {

	var fmf = []struct {
		strs     []string
		m, n     int
		expected int
	}{
		{[]string{"10", "0001", "111001", "1", "0"}, 5, 3, 4},
		{[]string{"10", "0", "1"}, 1, 1, 2},
	}

	for _, val := range fmf {
		actual := findMaxForm(val.strs, val.m, val.n)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
