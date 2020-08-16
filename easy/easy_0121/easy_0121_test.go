package easy_0121

import (
	"fmt"
	"testing"
)

func TestMaxProfit(t *testing.T) {
	var mp = []struct {
		prices   []int
		expected int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, val := range mp {
		actual := maxProfit(val.prices)

		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
