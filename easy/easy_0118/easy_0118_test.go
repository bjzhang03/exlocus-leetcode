package easy_0118

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {

	var ge = []struct {
		numRows  int
		expected [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
	}

	for _, val := range ge {
		actual := generate(val.numRows)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
