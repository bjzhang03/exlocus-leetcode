package medium_0388

import "testing"

func TestLengthLongestPath(t *testing.T) {

	var llp = []struct {
		in       string
		expected int
	}{
		{"dir\n\tsubdir1\n\t\tfile1.ext\n\t\tsubsubdir1\n\tsubdir2\n\t\tsubsubdir2\n\t\t\tfile2.ext", 32},
		{"a", 0},
	}

	for _, val := range llp {
		actual := lengthLongestPath(val.in)

		if actual != val.expected {
			t.Errorf("Test Failed! actual:= %d, expected:=%d", actual, val.expected)
		}
	}

}
