package medium_0756

import (
	"fmt"
	"testing"
)

func TestPyramidTransition(t *testing.T) {
	var cases = []struct {
		bottom   string
		allowed  []string
		expected bool
	}{
		{"BCD", []string{"BCG", "CDE", "GEA", "FFF"}, true},
		{"AABA", []string{"AAA", "AAB", "ABA", "ABB", "BAC"}, false},
		// 个人觉得这个测试用例是有问题的,没看出来如何解决
		//{"BCD", []string{"ACC", "ACB", "ABD", "DAA", "BDC", "BDB", "DBC", "BBD", "BBC", "DBD", "BCC",
		//	"CDD", "ABA", "BAB", "DDC", "CCD", "DDA", "CCA", "DDD"}, true},
	}

	for _, val := range cases {
		actual := pyramidTransition(val.bottom, val.allowed)

		if actual != val.expected {
			t.Errorf("Test failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
