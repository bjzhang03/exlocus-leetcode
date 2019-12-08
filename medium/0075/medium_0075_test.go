package _075

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {

	var sortColorsTest = []struct {
		in       []int // input
		expected []int // expected result
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, sct := range sortColorsTest {
		sortColors(sct.in)
		if !reflect.DeepEqual(sct.in, sct.expected) {
			t.Errorf("sortColors(%d) = %d; expected %d", sct.in, sct.in, sct.expected)
		}
	}

}
