package medium_0419

import "testing"

func TestCountBattleships(t *testing.T) {

	var cbs = []struct {
		board    [][]byte
		expected int
	}{
		{[][]byte{{'X', '.', '.', 'X'}, {'.', '.', '.', 'X'}, {'.', '.', '.', 'X'}}, 2},
		{[][]byte{{'.', '.', '.', 'X'}, {'X', 'X', 'X', 'X'}, {'.', '.', '.', 'X'}}, 0},
	}

	for _, val := range cbs {
		actual := countBattleships(val.board)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
