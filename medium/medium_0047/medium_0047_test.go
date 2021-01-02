package medium_0047

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermuteUnique(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{1, 1, 2}, [][]int{{1, 1, 2}, {1, 2, 1}, {2, 1, 1}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := permuteUnique(val.nums)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
