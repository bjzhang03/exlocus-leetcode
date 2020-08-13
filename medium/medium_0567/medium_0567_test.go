package medium_0567

import "testing"

func TestCheckInclusion(t *testing.T) {

	var ci = []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaoo", false},
		{"adc", "dcda", true},
	}

	for _, val := range ci {
		actual := checkInclusion(val.s1, val.s2)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
