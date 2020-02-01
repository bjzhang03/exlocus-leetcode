package medium_0386

import (
	"reflect"
	"testing"
)

func TestLexicalOrder(t *testing.T) {

	var lo = []struct {
		in       int
		expected []int
	}{
		{13, []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for _, val := range lo {
		actual := lexicalOrder(val.in)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Error("Test Failed!")
		}
	}

}
