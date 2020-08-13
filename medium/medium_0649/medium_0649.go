package medium_0649

func predictPartyVictory(senate string) string {
	if len(senate) > 0 {
		for len(senate) > 0 && !checkpass(senate) {
			twosenate := senate + senate
			// 找到第一个
			current := twosenate[0]
			for i := 0; i < len(twosenate); i++ {
				if twosenate[i] != current {
					twosenate = "" + twosenate[:i] + twosenate[i+1:]
					senate = twosenate[1:len(senate)]
					break
				}
			}
		}
		//判断最后的结果
		if len(senate) > 0 && senate[0] == 'R' {
			return "Radiant"

		} else if len(senate) > 0 && senate[0] == 'D' {
			return "Dire"
		}
	}
	return ""
}

func checkpass(senate string) bool {
	if len(senate) > 0 {
		current := senate[0]
		for i := 1; i < len(senate); i++ {
			if current != senate[i] {
				return false
			}
		}
	}
	return true
}
