package medium_0537

import "testing"

func TestComplexNumberMultiply(t *testing.T) {

	var cnm = []struct {
		a, b     string
		expected string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
	}

	for _, val := range cnm {
		if val.expected != complexNumberMultiply(val.a, val.b) {
			t.Error("Test Failed!")
		}
	}

}
