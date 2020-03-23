package medium_0650

import "testing"

func TestMinSteps(t *testing.T) {

	var ms = []struct {
		n        int
		expected int
	}{
		{3, 3},
		{10, 7},
		{19, 19},
		{14, 9},
		{100, 14},
	}

	for _, val := range ms {
		actual := minSteps(val.n)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %d, expected := %d", actual, val.expected)
		}
	}

}
