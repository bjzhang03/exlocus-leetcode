package medium_0649

import "testing"

func TestPredictPartyVictory(t *testing.T) {

	var ppv = []struct {
		senate   string
		expected string
	}{
		{"RD", "Radiant"},
		{"RDD", "Dire"},
		{"DDRRRR", "Radiant"},
	}

	for _, val := range ppv {
		actual := predictPartyVictory(val.senate)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
