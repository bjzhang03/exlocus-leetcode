package medium_0423

import (
	"math"
)

// https://blog.csdn.net/qq_36946274/article/details/81587007
func originalDigits(s string) string {
	result := ""
	if len(s) > 0 {
		// 统计各个字母出现的次数
		save := make(map[uint8]int, 26)
		for i := 0; i < len(s); i++ {
			if _, ok := save[s[i]]; !ok {
				save[s[i]] = 1
			} else {
				save[s[i]]++
			}
		}

		//numstable := map[string]uint8{"zero": '0', "one": '1', "two": '2', "three": '3', "four": '4', "five": '5', "six": '6', "seven": '7', "eight": '8', "nine": '9'}
		//numshort := map[string]string{"zero": "zero", "one": "one", "two": "two", "thre": "three", "four": "four", "five": "five", "six": "six", "sevn": "seven", "eight": "eight", "nie": "nine"}
		for hasStrs(save) {

		}

	}
	return result
}

// 找出当前的所有的排列数据

func deepFirstSearch(save []uint8, result []string) {

}

// 统计出最少使用的几个数字
func getUint8s(save map[uint8]int) []uint8 {
	result := []uint8{}
	if len(save) <= 5 {
		// 直接将uint8添加进来
		for key, _ := range save {
			result = append(result, key)
		}
	} else {
		//找到至少5个最少出现的uint8
		//这里的5是观察上面的数字特征得来的
		used := map[int]bool{}
		for len(result) < 5 {
			countmin := math.MaxInt32
			for _, val := range save {
				if _, ok := used[val]; !ok && countmin > val {
					countmin = val
				}
			}
			for key, val := range save {
				if countmin == val {
					result = append(result, key)
				}
			}
			used[countmin] = true
		}
	}
	return result
}

// 判断是否还有字符串可用
func hasStrs(save map[uint8]int) bool {
	for _, val := range save {
		if val > 0 {
			return true
		}
	}
	return false
}
