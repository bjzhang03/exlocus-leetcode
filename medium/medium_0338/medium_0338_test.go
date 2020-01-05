package medium_0338

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {

	var countBitTest = []struct {
		in       int
		expected []int
	}{
		{1, []int{0, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}

	for _, val := range countBitTest {
		result := countBits(val.in)

		if !reflect.DeepEqual(result, val.expected) {
			t.Errorf("test failed!")
		}
	}

}
