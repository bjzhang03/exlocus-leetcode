package medium_0049

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var ga = []struct {
		strs     []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
	}

	for _, val := range ga {
		actual := groupAnagrams(val.strs)

		if reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}

}
