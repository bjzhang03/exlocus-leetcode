package medium_0638

import (
	"fmt"
)

var (
	max_val = 100000000
)

func shoppingOffers(price []int, special [][]int, needs []int) int {
	if checkpass(needs) {
		return solve(&price, &special, needs, &map[string]int{fmt.Sprint(make([]int, len(needs))): 0})
	}
	return 0
}

func solve(price *[]int, special *[][]int, needs []int, used *map[string]int) int {
	needsstr := fmt.Sprint(needs)
	if _, ok := (*used)[needsstr]; ok {
		return (*used)[needsstr]
	} else if !ok {
		result := max_val
		// 先查看优惠大礼包
		for i := 0; i < len(*special); i++ {
			// 构造下一次的需求
			for j := 0; j < len(needs); j++ {
				needs[j] = needs[j] - (*special)[i][j]
			}
			// 先进行checkpass的判断,少了一些堆栈调用
			if checkpass(needs) {
				if solveresult := solve(price, special, needs, used); result > (*special)[i][len(needs)]+solveresult {
					result = (*special)[i][len(needs)] + solveresult
				}
			}
			// 恢复现场
			for j := 0; j < len(needs); j++ {
				needs[j] = needs[j] + (*special)[i][j]
			}
		}
		// 再看单个购买
		for i := 0; i < len(*price); i++ {
			needs[i] = needs[i] - 1
			if checkpass(needs) {
				if solveresult := solve(price, special, needs, used); result > (*price)[i]+solveresult {
					result = (*price)[i] + solveresult
				}
			}
			needs[i] = needs[i] + 1
		}
		(*used)[needsstr] = result
		return result
	}
	return max_val
}

func checkpass(needs []int) bool {
	for i := 0; i < len(needs); i++ {
		if needs[i] < 0 {
			return false
		}
	}
	return true
}
