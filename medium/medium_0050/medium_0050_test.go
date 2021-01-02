package medium_0050

import (
	"fmt"
	"math"
	"testing"
)

func TestMyPow(t *testing.T) {
	var cases = []struct {
		x        float64
		n        int
		expected float64
	}{
		{2.00000, 10, 1024.00000},
		{2.10000, 3, 9.26100},
		{2.00000, -2, 0.25000},
	}

	for _, val := range cases {
		actual := myPow(val.x, val.n)

		if math.Abs(actual-val.expected) > 0.0001 {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
