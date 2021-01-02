package medium_0049

import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	var cases = []struct {
		strs     []string
		expected [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"ate", "eat", "tea"}, {"nat", "tan"}, {"bat"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}

	for _, val := range cases {
		actual := groupAnagrams(val.strs)

		if !stringSetEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual := %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}

func stringSetEqual(a [][]string, b [][]string) bool {
	// 长度相等
	if len(a) == len(b) {
		//判断集合中是否含有某个元素的函数
		setcontainsf := func(sets [][]string, item []string) bool {
			for j := 0; j < len(sets); j++ {
				sort.Strings(sets[j])
				if reflect.DeepEqual(sets[j], item) {
					//fmt.Printf("item := %s, sets := %s", fmt.Sprint(item), fmt.Sprint(sets[j]))
					return true
				}
			}
			return false
		}
		// 每一个a中的元素都在b中出现
		for i := 0; i < len(a); i++ {
			sort.Strings(a[i])
			// 出现集合中没有
			if !setcontainsf(b, a[i]) {
				return false
			}
		}
		//每一个b中的元素都在a中出现
		for i := 0; i < len(b); i++ {
			sort.Strings(b[i])
			if !setcontainsf(a, b[i]) {
				return false
			}
		}
		return true
	}
	return false
}
