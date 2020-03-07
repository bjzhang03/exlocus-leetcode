package medium_0526

import "testing"

func TestCountArrangement(t *testing.T) {

	var ca = []struct {
		N        int
		expected int
	}{
		{2, 2},
		{10, 700},
		{15, 24679},
	}

	for _, val := range ca {
		if val.expected != countArrangement(val.N) {
			t.Error("Test Failed!")
		}
	}

}
