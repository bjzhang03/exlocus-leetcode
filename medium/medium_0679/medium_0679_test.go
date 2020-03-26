package medium_0679

import (
	"fmt"
	"testing"
)

func TestJudgePoint24(t *testing.T) {
	var jp = []struct {
		nums     []int
		expected bool
	}{
		{[]int{4, 1, 8, 7}, true},
		{[]int{1, 2, 1, 2}, false},
		{[]int{1, 5, 9, 1}, false},
		{[]int{3, 3, 8, 8}, true},
	}

	fmt.Println(deepFirstSearch([]float64{8, 3, 8, 3}))

	for _, val := range jp {
		actual := judgePoint24(val.nums)
		if actual != val.expected {
			t.Errorf("Test Failed! actual := %t , expected := %t", actual, val.expected)
		}
	}

}
