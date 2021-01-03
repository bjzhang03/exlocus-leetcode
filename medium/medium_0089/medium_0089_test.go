package medium_0089

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGrayCode(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		n        int
		expected []int
	}{
		{3, []int{0, 4, 6, 2, 3, 7, 5, 1}},
		{2, []int{0, 1, 3, 2}},
		{0, []int{0}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		actual := grayCode(val.n)
		// 先对数据进行排序,方便比较
		sort.Ints(actual)
		/*验证结果*/
		if sort.Ints(val.expected); !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
