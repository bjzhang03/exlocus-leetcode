package medium_0064

import (
	"fmt"
	"testing"
)

func TestMinPathSum(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		grid     [][]int
		expected int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := minPathSum(val.grid)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
