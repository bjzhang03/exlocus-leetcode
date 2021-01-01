package medium_0018

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		target   int
		expected [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{}, 0, [][]int{}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := fourSum(val.nums, val.target)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
