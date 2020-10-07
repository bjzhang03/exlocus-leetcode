package easy_0226

import (
	"fmt"
	"testing"
)

func TestInvertTree(t *testing.T) {

	var cases = []struct {
		root     *TreeNode
		expected *TreeNode
	}{
		{buildTreeFromSlice([]int{4, 2, 7, 1, 3, 6, 9}, 0), buildTreeFromSlice([]int{4, 7, 2, 9, 6, 3, 1}, 0)},
	}

	for _, val := range cases {
		actual := invertTree(val.root)
		// fmt.Println(reflect.DeepEqual(actual, val.expected))
		if !actual.Equal(val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
