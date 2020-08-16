package easy_0119

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {

	var gr = []struct {
		rowIndex int
		expected []int
	}{
		{3, []int{1, 3, 3, 1}},
		{4, []int{1, 4, 6, 4, 1}},
	}

	for _, val := range gr {
		actual := getRow(val.rowIndex)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))

		}
	}

}
