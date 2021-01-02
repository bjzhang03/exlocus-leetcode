package medium_0039

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}},
		{[]int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}},
		{[]int{2}, 1, [][]int{}},
		{[]int{1}, 1, [][]int{{1}}},
		{[]int{1}, 2, [][]int{{1, 1}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := combinationSum(val.candidates, val.target)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
