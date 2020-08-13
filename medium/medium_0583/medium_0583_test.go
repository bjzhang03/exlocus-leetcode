package medium_0583

import "testing"

func TestMinDistance(t *testing.T) {
	var md = []struct {
		word1    string
		word2    string
		expected int
	}{
		{"sea", "eat", 2},
	}

	for _, val := range md {
		actual := minDistance(val.word1, val.word2)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
