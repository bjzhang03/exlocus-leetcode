package easy_0066

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPlusOne(t *testing.T) {
	var po = []struct {
		digits   []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
	}

	for _, val := range po {
		actual := plusOne(val.digits)
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))

		}
	}

}
