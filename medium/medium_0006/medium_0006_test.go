package medium_0006

import (
	"fmt"
	"testing"
)

func TestConvert(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		s        string
		numRows  int
		expected string
	}{
		{"PAYPALISHIRING", 3, "PAHNAPLSIIGYIR"},
		{"PAYPALISHIRING", 4, "PINALSIGYAHRPI"},
		{"A", 1, "A"},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := convert(val.s, val.numRows)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
