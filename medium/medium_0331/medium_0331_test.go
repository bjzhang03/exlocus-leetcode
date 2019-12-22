package medium_0331

import "testing"

func TestIsValidSerialization(t *testing.T) {

	var validser = []struct {
		in       string
		expected bool
	}{
		{"9,3,4,#,#,1,#,#,2,#,6,#,#", true},
		{"9,#,#,1", false},
	}

	for _, vas := range validser {
		actual := isValidSerialization(vas.in)

		if actual != vas.expected {
			t.Errorf("test error!")
		}
	}

}
