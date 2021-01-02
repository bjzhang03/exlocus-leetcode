package medium_0046

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPermute(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}},
		{[]int{0, 1}, [][]int{{0, 1}, {1, 0}}},
		{[]int{1}, [][]int{{1}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := permute(val.nums)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
