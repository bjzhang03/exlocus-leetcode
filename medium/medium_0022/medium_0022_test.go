package medium_0022

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		n        int
		expected []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := generateParenthesis(val.n)
		/*验证结果*/
		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
