package medium_0319

import "testing"

func TestBulbSwitch(t *testing.T) {

	var bulbTests = []struct {
		in       int
		expected int
	}{
		//{20, 2},
	}

	for _, bt := range bulbTests {
		actual := bulbSwitch(bt.in)

		if actual != bt.expected {
			t.Errorf("Test failed!")
		}
	}

}
