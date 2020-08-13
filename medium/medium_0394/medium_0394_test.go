package medium_0394

import "testing"

func TestDecodeString(t *testing.T) {

	var ds = []struct {
		in       string
		expected string
	}{
		{"3[a2[c]]", "accaccacc"},
		{"3[a]2[bc]", "aaabcbc"},
		{"2[abc]3[cd]ef", "abcabccdcdcdef"},
	}

	for _, val := range ds {
		actual := decodeString(val.in)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %s, expected:= %s", actual, val.expected)
		}
	}

}
