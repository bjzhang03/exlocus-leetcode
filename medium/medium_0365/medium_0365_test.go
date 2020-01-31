package medium_0365

import "testing"

type mwater struct {
	x, y, z int
}

func TestCanMeasureWater(t *testing.T) {

	var cmws = []struct {
		in       mwater
		expected bool
	}{
		{mwater{1, 1, 12}, false},
	}

	for _, val := range cmws {
		actual := canMeasureWater(val.in.x, val.in.y, val.in.z)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
