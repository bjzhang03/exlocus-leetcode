package easy_0001

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{}, 0, []int{0, 0}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := twoSum(val.nums, val.target)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
