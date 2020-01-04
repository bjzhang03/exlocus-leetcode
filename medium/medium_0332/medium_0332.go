package medium_0332

func deepFirstSearch(tickets [][]string, save []string, result *[]string) {
	if len(tickets) == 0 {
		if len(*result) == 0 {
			*result = append(*result, save...)
		} else {
			*result = append([]string{}, maxItem(*result, save)...)
		}
	} else if len(tickets) > 0 {
		for i := 0; i < len(tickets); i++ {
			// 我在这里是直接全部新建数据
			if tickets[i][0] == save[len(save)-1] {
				nexttickets := append(append([][]string{}, tickets[:i]...), tickets[i+1:]...)
				save = append(save, tickets[i][1])
				deepFirstSearch(nexttickets, save, result)
				save = save[:len(save)-1]
			}
		}

	}
}

func maxItem(itema []string, itemb []string) []string {
	for i := 0; i < len(itema); i++ {
		if itema[i] > itemb[i] {
			// 大于的话,则替换
			return itemb
		} else if itema[i] < itemb[i] {
			// 小于的话则退出
			return itema
		}
		// 等于则进行下一步判断
	}
	return itema
}

func findItinerary(tickets [][]string) []string {
	if len(tickets) > 0 {
		resultc := []string{}
		deepFirstSearch(tickets, []string{"JFK"}, &resultc)

		return resultc
	}
	return nil
}
