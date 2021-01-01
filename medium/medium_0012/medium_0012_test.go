package medium_0012

import (
	"fmt"
	"testing"
)

func TestIntToRoman(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		num      int
		expected string
	}{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := intToRoman(val.num)
		/*验证结果*/
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
