package medium_0332

import (
	"reflect"
	"testing"
)

func TestFindItinerary(t *testing.T) {
	var find = []struct {
		in       [][]string
		expected []string
	}{
		//{[][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{[][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
	}

	for _, f := range find {
		actual := findItinerary(f.in)

		if !reflect.DeepEqual(actual, f.expected) {
			t.Errorf("test failed!")

		}
	}
}
