package medium_0079

import (
	"fmt"
	"testing"
)

func TestExist(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		board    [][]byte
		word     string
		expected bool
	}{
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE", true},
		{[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB", false},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := exist(val.board, val.word)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
