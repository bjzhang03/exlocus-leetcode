package medium_0672

import "testing"

func TestFlipLights(t *testing.T) {

	var fl = []struct {
		n, m     int
		expected int
	}{
		{1, 1, 2},
		{2, 1, 3},
		{3, 1, 4},
		{1, 0, 1},
		{1000, 999, 8},
		{3, 6, 8},
	}

	for _, val := range fl {
		actual := flipLights(val.n, val.m)

		if actual != val.expected {
			t.Errorf("Test Failed! actual:= %d , expected := %d", actual, val.expected)
		}
	}

}
