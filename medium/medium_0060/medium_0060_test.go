package medium_0060

import (
	"reflect"
	"testing"
)

func TestGetPermutation(t *testing.T) {

	type perm struct {
		n int
		k int
	}

	var permTest = []struct {
		in       perm
		expected string
	}{
		//{perm{3, 3}, "213"},
		//{perm{4, 9}, "2314"},
		//{perm{3, 2}, "132"},
		//{perm{3, 1}, "123"},
		//{perm{2, 2}, "21"},
		//{perm{4, 2}, "1243"},
	}

	for _, pt := range permTest {
		actual := getPermutation(pt.in.n, pt.in.k)
		if !reflect.DeepEqual(actual, pt.expected) {
			t.Errorf("actual = %s, expected =%s", actual, pt.expected)
		}
	}

}
