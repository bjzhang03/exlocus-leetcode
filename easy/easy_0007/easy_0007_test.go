package easy_0007

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	/*构建矩阵测试的结构*/
	var rev = []struct {
		x        int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}

	for _, val := range rev {
		actual := reverse(val.x)
		if actual != val.expected {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
