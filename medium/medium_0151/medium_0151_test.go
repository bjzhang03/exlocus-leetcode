package medium_0151

import "testing"

func TestReverseWords(t *testing.T) {

	var reverseTests = []struct {
		in       string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world!  ", "world! hello"},
		{"a good   example", "example good a"},
	}

	for _, rt := range reverseTests {
		actual := reverseWords(rt.in)

		if actual != rt.expected {
			t.Errorf("actual = %s, expected = %s", actual, rt.expected)
		}
	}

}
