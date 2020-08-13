package medium_0433

import "testing"

func TestMinMutation(t *testing.T) {

	var mms = []struct {
		start    string
		end      string
		bank     []string
		expected int
	}{
		{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}, 1},
		{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}, 2},
		{"AAAAACCC", "AACCCCCC", []string{"AAAACCCC", "AAACCCCC", "AACCCCCC"}, 3},
		{"AACCGGTT", "AACCGGTA", []string{}, -1},
	}

	for _, val := range mms {
		actual := minMutation(val.start, val.end, val.bank)

		if actual != val.expected {
			t.Errorf("Test Failed! actual:= %d , expected := %d", actual, val.expected)
		}
	}

}
