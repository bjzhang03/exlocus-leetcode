package medium_0739

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {

	var dt = []struct {
		T        []int
		expected []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}

	for _, val := range dt {
		actual := dailyTemperatures(val.T)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! input := %v, expected := %v, actual := %v", val.T, val.expected, actual)

		}
	}

}
