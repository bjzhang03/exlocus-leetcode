package medium_0043

import (
	"fmt"
	"testing"
)

func TestMultiply(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		num1     string
		num2     string
		expected string
	}{
		{"2", "3", "6"},
		{"123", "456", "56088"},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := multiply(val.num1, val.num2)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
