package medium_0034

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchRange(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
		{[]int{}, 0, []int{-1, -1}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := searchRange(val.nums, val.target)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
