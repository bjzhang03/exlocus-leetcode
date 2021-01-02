package medium_0040

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombinationSum2(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		candidates []int
		target     int
		expected   [][]int
	}{
		{[]int{10, 1, 2, 7, 6, 1, 5}, 8, [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}},
		{[]int{2, 5, 2, 1, 2}, 5, [][]int{{1, 2, 2}, {5}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := combinationSum2(val.candidates, val.target)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
