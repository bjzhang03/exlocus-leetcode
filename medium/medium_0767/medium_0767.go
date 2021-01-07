package medium_0767

import (
	"fmt"
	"math"
)

func reorganizeString(S string) string {
	return solve(S)
}

func solve(str string) string {
	result := ""
	// 先统计每一个字符串出现的字数
	save := map[uint8]int{}
	for i := 0; i < len(str); i++ {
		if _, ok := save[str[i]]; !ok {
			save[str[i]] = 1
		} else if ok {
			save[str[i]]++
		}
	}
	for len(save) > 0 {
		// 找到save中最大的和第二大的两个数字
		fsig := uint8('0')
		fmax := math.MinInt32
		ssig := uint8('0')
		smax := math.MinInt32

		// 从数组中找出出现最多,第二多的数据
		for key, val := range save {
			if fmax == math.MinInt32 {
				fmax = val
				fsig = key
			} else if smax == math.MinInt32 {
				if fmax > val {
					smax = val
					ssig = key
				} else {
					smax = fmax
					ssig = fsig
					fmax = val
					fsig = key
				}
			} else {
				if smax < val {
					if fmax < val {
						smax = fmax
						ssig = fsig
						fmax = val
						fsig = key
					} else {
						smax = val
						ssig = key
					}
				}
			}
		}
		if fmax == smax && len(result) > 0 {
			if result[len(result)-1] == fsig {
				if ssig != uint8('0') {
					result = result + fmt.Sprintf("%c", ssig)
					save[ssig]--
					if save[ssig] == 0 {
						delete(save, ssig)
					}
				}
				if fsig != uint8('0') {
					result = result + fmt.Sprintf("%c", fsig)
					save[fsig]--
					if save[fsig] == 0 {
						delete(save, fsig)
					}
				}
			} else {
				if fsig != uint8('0') {
					result = result + fmt.Sprintf("%c", fsig)
					save[fsig]--
					if save[fsig] == 0 {
						delete(save, fsig)
					}
				}
				if ssig != uint8('0') {
					result = result + fmt.Sprintf("%c", ssig)
					save[ssig]--
					if save[ssig] == 0 {
						delete(save, ssig)
					}
				}
			}
		} else {
			if fsig != uint8('0') {
				result = result + fmt.Sprintf("%c", fsig)
				save[fsig]--
				if save[fsig] == 0 {
					delete(save, fsig)
				}
			}
			if ssig != uint8('0') {
				result = result + fmt.Sprintf("%c", ssig)
				save[ssig]--
				if save[ssig] == 0 {
					delete(save, ssig)
				}
			}
		}

	}

	if !check(result) {
		return ""
	}
	return result
}

func check(str string) bool {
	for i := 0; i < len(str); i++ {
		if i+1 < len(str) && str[i] == str[i+1] {
			return false
		}
	}
	return true
}
