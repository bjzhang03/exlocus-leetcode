package medium_0481

import "testing"

func TestMagicalString(t *testing.T) {

	var ms = []struct {
		n        int
		expected int
	}{
		{6, 3},
		{20, 10},
		{4, 2},
		{7, 4},
	}

	for _, val := range ms {
		actual := magicalString(val.n)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
