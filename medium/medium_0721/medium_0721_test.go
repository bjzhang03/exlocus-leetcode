package medium_0721

import (
	"reflect"
	"testing"
)

func TestAccountsMerge(t *testing.T) {
	var am = []struct {
		accounts [][]string
		expected [][]string
	}{
		{[][]string{{"John", "johnsmith@mail.com", "john00@mail.com"}, {"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}},
			[][]string{{"John", "john00@mail.com", "john_newyork@mail.com", "johnsmith@mail.com"}, {"John", "johnnybravo@mail.com"}, {"Mary", "mary@mail.com"}}},
	}

	for _, val := range am {
		actual := accountsMerge(val.accounts)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! actual := %s,expected := %s\n", actual, val.expected)
		}
	}
}
