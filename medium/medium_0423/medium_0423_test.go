package medium_0423

import "testing"

func TestOriginalDigits(t *testing.T) {

	var ods = []struct {
		in       string
		expected string
	}{
		{"owoztneoer", "012"},
	}

	for _, val := range ods {
		actual := originalDigits(val.in)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
