package medium_0073

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetZeroes(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		matrix   [][]int
		expected [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}},
		{[][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}, [][]int{{0, 0, 0, 0}, {0, 4, 5, 0}, {0, 3, 1, 0}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		setZeroes(val.matrix)
		/*验证结果*/
		if !reflect.DeepEqual(val.matrix, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.matrix))
		}
	}
}
