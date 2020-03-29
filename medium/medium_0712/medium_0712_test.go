package medium_0712

import "testing"

func TestMinimumDeleteSum(t *testing.T) {
	var mds = []struct {
		s1       string
		s2       string
		expected int
	}{
		{"sea", "eat", 231},
		{"delete", "leet", 403},
		{"ihlnqpdwqgcd", "mgrumwmpjedv", 2165},
	}

	for _, val := range mds {
		actual := minimumDeleteSum(val.s1, val.s2)

		if actual != val.expected {
			t.Errorf("Test Failed!,actual := %d ,expected := %d", actual, val.expected)
		}
	}

}
