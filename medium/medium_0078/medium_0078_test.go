package medium_0078

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSubsets(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		nums     []int
		expected [][]int
	}{
		{[]int{1, 2, 3}, [][]int{{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3}}},
		{[]int{0}, [][]int{{}, {0}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := subsets(val.nums)
		/*验证结果*/
		if !sliceSetEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}

/*判断集合是否相等*/
func sliceSetEqual(a [][]int, b [][]int) bool {
	if len(a) == len(b) {
		setcontainsf := func(set [][]int, item []int) bool {
			for i := 0; i < len(set); i++ {
				if reflect.DeepEqual(set[i], item) {
					return true
				}
			}
			return false
		}
		/*判断所有的a都在b中*/
		for i := 0; i < len(a); i++ {
			if !setcontainsf(b, a[i]) {
				return false
			}
		}
		/*判断所有的b都在a中*/
		for i := 0; i < len(b); i++ {
			if !setcontainsf(a, b[i]) {
				return false
			}
		}
		return true
	}
	return false
}
