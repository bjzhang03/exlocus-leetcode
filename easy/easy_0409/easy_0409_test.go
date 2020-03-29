package easy_0409

import "testing"

func TestLongestPalindrome(t *testing.T) {
	var lp = []struct {
		s        string
		expected int
	}{
		{"abccccdd", 7},
		{"civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthat" +
			"fieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewec" +
			"annotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenot" +
			"lenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfar" +
			"sonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethela" +
			"stpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeopl" +
			"ebythepeopleforthepeopleshallnotperishfromtheearth", 983},
	}

	for _, val := range lp {
		actual := longestPalindrome(val.s)

		if actual != val.expected {
			t.Errorf("Test Failed! actual := %d, expected := %d", actual, val.expected)
		}
	}
}
