package medium_0029

import (
	"testing"
)

func TestDivide(t *testing.T) {

	type divideStr struct {
		dividend int
		divisor  int
	}

	var divideTest = []struct {
		in       divideStr // input
		expected int       // expected result
	}{
		{divideStr{10, 3}, 3},
		{divideStr{7, -3}, -2},
		{divideStr{-2147483648, -1}, 2147483647},
	}

	for _, dt := range divideTest {
		actual := divide(dt.in.dividend, dt.in.divisor)
		if dt.expected != actual {
			t.Errorf("actual = %d; expected %d", actual, dt.expected)
		}
	}

}
