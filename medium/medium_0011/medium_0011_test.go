package medium_0011

import (
	"fmt"
	"testing"
)

func TestMaxArea(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		height   []int
		expected int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := maxArea(val.height)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
