package medium_0647

import "testing"

func TestCountSubstrings(t *testing.T) {

	var cs = []struct {
		s        string
		expected int
	}{
		{"abc", 3},
		{"aaa", 6},
	}

	for _, val := range cs {
		actual := countSubstrings(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
