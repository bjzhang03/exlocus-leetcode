package medium_0357

import (
	"fmt"
	"testing"
)

func TestCountNumbersWithUniqueDigits(t *testing.T) {

	var cnwus = []struct {
		in       int
		expected int
	}{
		{2, 91},
		{3, 739},
		{4, 5275},
	}

	fmt.Println(allsort(1, 1))

	fmt.Println(allselect(10, 2))

	for _, value := range cnwus {
		actual := countNumbersWithUniqueDigits(value.in)

		if actual != value.expected {
			t.Errorf("Test Error!")
		}
	}

}
