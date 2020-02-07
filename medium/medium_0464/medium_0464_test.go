package medium_0464

import "testing"

func TestCanIWin(t *testing.T) {

	var ciw = []struct {
		maxChoosableInteger int
		desiredTotal        int
		expected            bool
	}{
		//{10, 11, false},
		//{10, 20, true},
		{19, 100, true},
		//{10, 45, true},
		//{10, 0, true},
		//{10, 40, false},
		//{18, 79, true},
	}

	for _, val := range ciw {
		actual := canIWin(val.maxChoosableInteger, val.desiredTotal)

		if actual != val.expected {
			t.Error("Test Failed!", val.maxChoosableInteger, val.desiredTotal, val.expected, actual)
		}
	}

}
