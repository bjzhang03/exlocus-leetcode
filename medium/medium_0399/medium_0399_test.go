package medium_0399

import (
	"reflect"
	"testing"
)

func TestCalcEquation(t *testing.T) {

	var ce = []struct {
		equations [][]string
		values    []float64
		queries   [][]string
		expected  []float64
	}{
		{[][]string{{"a", "b"}, {"b", "c"}}, []float64{2, 3}, [][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}, []float64{6, 0.5, -1, 1, -1}},
		{[][]string{{"x1", "x2"}, {"x2", "x3"}, {"x3", "x4"}, {"x4", "x5"}}, []float64{3, 4, 5, 6}, [][]string{{"x1", "x5"}, {"x5", "x2"}, {"x2", "x4"}, {"x2", "x2"}, {"x2", "x9"}, {"x9", "x9"}}, []float64{360, 0.008333333333333333, 20, 1, -1, -1}},
	}

	for _, val := range ce {
		actual := calcEquation(val.equations, val.values, val.queries)

		if !reflect.DeepEqual(actual, val.expected) {
			t.Error("Test Failed!")
		}
	}

}
