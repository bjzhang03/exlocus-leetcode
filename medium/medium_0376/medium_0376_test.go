package medium_0376

import "testing"

func TestWiggleMaxLength(t *testing.T) {

	var wml = []struct {
		in       []int
		expected int
	}{

		{[]int{1, 7, 4, 9, 2, 5}, 6},
	}

	for _, val := range wml {
		actual := wiggleMaxLength(val.in)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
