package medium_0016

import (
	"fmt"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{-1, 2, 1, -4}, 1, 2},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := threeSumClosest(val.nums, val.target)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
