package medium_0357

import (
	"testing"
)

func TestCountNumbersWithUniqueDigits(t *testing.T) {

	var cnwus = []struct {
		in       int
		expected int
	}{
		//{2, 91},
		//{3, 739},
		//{4, 5275},
	}

	for _, value := range cnwus {
		actual := countNumbersWithUniqueDigits(value.in)

		if actual != value.expected {
			t.Errorf("Test Error!")
		}
	}

}
