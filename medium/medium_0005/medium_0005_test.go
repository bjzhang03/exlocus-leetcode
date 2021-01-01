package medium_0005

import (
	"fmt"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		s        string
		expected string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := longestPalindrome(val.s)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
