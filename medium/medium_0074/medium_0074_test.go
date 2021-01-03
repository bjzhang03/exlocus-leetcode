package medium_0074

import (
	"fmt"
	"testing"
)

func TestSearchMatrix(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		matrix   [][]int
		target   int
		expected bool
	}{
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3, true},
		{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13, false},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := searchMatrix(val.matrix, val.target)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
