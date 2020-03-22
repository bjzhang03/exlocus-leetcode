package medium_0648

import "testing"

func TestReplaceWords(t *testing.T) {
	var rw = []struct {
		dict     []string
		sentence string
		expected string
	}{
		{[]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery", "the cat was rat by the bat"},
		{[]string{"a", "aa", "aaa", "aaaa"}, "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa", "a a a a a a a a bbb baba a"},
	}

	for _, val := range rw {
		actual := replaceWords(val.dict, val.sentence)

		if actual != val.expected {
			t.Errorf("Test Failed!")
		}
	}

}
