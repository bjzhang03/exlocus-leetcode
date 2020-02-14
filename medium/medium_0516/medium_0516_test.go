package medium_0516

import "testing"

func TestLongestPalindromeSubseq(t *testing.T) {

	var lps = []struct {
		s        string
		expected int
	}{
		{"bbbab", 4},
		{"cbbd", 2},
		{"cbbdabab", 5},
		{"abcdef", 1},
		{"euazbipzncptldueeuechubrcourfpftcebikrxhybkymimgvldiwqvkszfycvqyvtiwfckexmowcxztkfyzqovbtmzpxojfofbvwnncajvrvdbvjhcrameamcfmcoxryjukhpljwszknhiypvyskmsujkuggpztltpgoczafmfelahqwjbhxtjmebnymdyxoeodqmvkxittxjnlltmoobsgzdfhismogqfpfhvqnxeuosjqqalvwhsidgiavcatjjgeztrjuoixxxoznklcxolgpuktirmduxdywwlbikaqkqajzbsjvdgjcnbtfksqhquiwnwflkldgdrqrnwmshdpykicozfowmumzeuznolmgjlltypyufpzjpuvucmesnnrwppheizkapovoloneaxpfinaontwtdqsdvzmqlgkdxlbeguackbdkftzbnynmcejtwudocemcfnuzbttcoew", 159},
	}

	for _, val := range lps {

		actual := longestPalindromeSubseq(val.s)

		if actual != val.expected {
			t.Error("Test Failed!")
		}
	}

}
