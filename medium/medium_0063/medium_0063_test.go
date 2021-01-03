package medium_0063

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		obstacleGrid [][]int
		expected     int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := uniquePathsWithObstacles(val.obstacleGrid)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
