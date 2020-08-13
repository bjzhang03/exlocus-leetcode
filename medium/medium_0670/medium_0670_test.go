package medium_0670

import "testing"

func TestMaximumSwap(t *testing.T) {
	var ms = []struct {
		num      int
		expected int
	}{
		{2736, 7236},
		{9973, 9973},
		{1993, 9913},
	}

	for _, val := range ms {
		actual := maximumSwap(val.num)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %d ,expected := %d", actual, val.expected)
		}
	}

}
