package medium_0062

import (
	"fmt"
	"testing"
)

func TestUniquePaths(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		m        int
		n        int
		expected int
	}{
		{3, 7, 28},
		{3, 2, 3},
		{7, 3, 28},
		{3, 3, 6},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := uniquePaths(val.m, val.n)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
