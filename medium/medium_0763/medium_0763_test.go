package medium_0763

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPartitionLabels(t *testing.T) {
	var cases = []struct {
		S        string
		expected []int
	}{
		{"ababcbacadefegdehijhklij", []int{9, 7, 8}},
	}

	for _, val := range cases {
		actual := partitionLabels(val.S)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Errorf("Test Failed! expected := %s, actual:= %s", fmt.Sprint(val.expected), fmt.Sprint(actual))
		}
	}
}
