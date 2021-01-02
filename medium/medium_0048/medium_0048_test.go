package medium_0048

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	/*构建矩阵测试的结构*/
	var cases = []struct {
		matrix   [][]int
		expected [][]int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}}},
		{[][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}},
			[][]int{{15, 13, 2, 5}, {14, 3, 4, 1}, {12, 6, 8, 9}, {16, 7, 10, 11}}},
		{[][]int{{1}}, [][]int{{1}}},
		{[][]int{{1, 2}, {3, 4}}, [][]int{{3, 1}, {4, 2}}},
	}
	/*执行矩阵测试*/
	for _, val := range cases {
		rotate(val.matrix)
		/*验证结果*/
		if !reflect.DeepEqual(val.matrix, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(val.matrix))
		}
	}
}
