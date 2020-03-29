package medium_0688

import (
	"math"
	"testing"
)

func TestKnightProbability(t *testing.T) {

	var kp = []struct {
		N        int
		K        int
		r        int
		c        int
		expected float64
	}{
		{3, 2, 0, 0, 0.0625},
		{1, 0, 0, 0, 1},
		{3, 2, 1, 1, 0},
		{8, 30, 6, 4, 0.00019},
		// 这个测试用例存在精度损失的问题
	}

	for _, val := range kp {
		actual := knightProbability(val.N, val.K, val.r, val.c)

		if math.Abs(actual-val.expected) > 0.000001 {
			t.Errorf("Test Failed! actual :=%f ,expected := %f", actual, val.expected)
		}
	}

}
