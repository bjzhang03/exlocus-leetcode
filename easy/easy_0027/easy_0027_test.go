package easy_0027

import (
	"fmt"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	var re = []struct {
		nums     []int
		val      int
		expected int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, val := range re {
		actual := removeElement(val.nums, val.val)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
