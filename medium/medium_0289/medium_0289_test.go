package medium_0289

import (
	"reflect"
	"testing"
)

func TestGameOfLife(t *testing.T) {

	var gameOfLifeTest = []struct {
		in       [][]int
		expected [][]int
	}{
		{[][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0,}}, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}},
	}

	for _, golt := range gameOfLifeTest {
		gameOfLife(golt.in)
		if reflect.DeepEqual(golt.in, golt.expected) {
			t.Errorf("Test Failed")
		}
	}

}
