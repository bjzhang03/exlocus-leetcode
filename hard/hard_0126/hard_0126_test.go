package hard_0126

import (
	"reflect"
	"testing"
)

func TestFindLadders(t *testing.T) {
	type wordLadder struct {
		beginWord string
		endWord   string
		wordList  []string
	}

	var ladderTest = []struct {
		in       wordLadder // input
		expected [][]string // expected result
	}{
		{wordLadder{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log", "cog"}}, [][]string{{"hit", "hot", "dot", "dog", "cog"}, {"hit", "hot", "lot", "log", "cog"}}},
		//{wordLadder{"hit", "cog", []string{"hot", "dot", "dog", "lot", "log"}}, nil},
		//{wordLadder{"hot", "dog", []string{"hot", "dog", "dot"}}, nil},
		{wordLadder{"qa", "sq", []string{"si", "go", "se", "cm", "so", "ph", "mt", "db", "mb",
			"sb", "kr", "ln", "tm", "le", "av", "sm", "ar", "ci", "ca", "br", "ti", "ba", "to", "ra", "fa", "yo", "ow", "sn", "ya",
			"cr", "po", "fe", "ho", "ma", "re", "or", "rn", "au", "ur", "rh", "sr", "tc", "lt", "lo", "as", "fr", "nb", "yb", "if",
			"pb", "ge", "th", "pm", "rb", "sh", "co", "ga", "li", "ha", "hz", "no", "bi", "di", "hi", "qa", "pi", "os", "uh", "wm",
			"an", "me", "mo", "na", "la", "st", "er", "sc", "ne", "mn", "mi", "am", "ex", "pt", "io", "be", "fm", "ta", "tb", "ni",
			"mr", "pa", "he", "lr", "sq", "ye"}}, nil},
	}

	for _, lt := range ladderTest {
		actual := findLadders(lt.in.beginWord, lt.in.endWord, lt.in.wordList)
		if !reflect.DeepEqual(actual, lt.expected) {
			t.Errorf("ladderLength(%s) = %s; expected %s", lt.in, actual, lt.expected)
		}
	}

}
