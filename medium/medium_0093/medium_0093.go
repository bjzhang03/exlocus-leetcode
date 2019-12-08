package medium_0093

import (
	"strconv"
)

// 将ip地址数据转化为string格式的数据
func ipAddressToStr(nums []int) string {
	result := ""
	for i := 0; i < len(nums)-1; i++ {
		result = result + strconv.Itoa(nums[i]) + "."
	}
	result = result + strconv.Itoa(nums[len(nums)-1])
	return result
}

// 判断是否是ip地址的操作
func isIpAddress(nums []int) bool {
	for _, key := range nums {
		if key > 255 || key < 0 {
			return false
		}
	}
	return true
}

// 深度优先回溯找到所有的ip地址组合
func dfs(result *[]string, tmp []int, index int, str string, usedSet *map[string]bool) {
	if (len(tmp) == 4 && index == len(str)) {
		if isIpAddress(tmp) {
			ipStr := ipAddressToStr(tmp)
			// 除了判断是否已经使用过以外,还要判断字符的长度是否正确
			if ok := (*usedSet)[ipStr]; !ok && len(ipStr) == len(str)+3 {
				*result = append(*result, ipStr)
			}
		}
	} else if index > len(str) || (len(tmp) >= 4 && index != len(str)) {
		return
	} else {
		for i := 1; i <= 3; i++ {
			// 找出将要处理的数据
			mystr := ""
			if index+i <= len(str) {
				mystr = str[index : index+i]
			} else {
				mystr = str[index:len(str)]
			}

			myint, err := strconv.Atoi(mystr)
			if err != nil {
				return
			}
			tmp = append(tmp, myint)
			dfs(result, tmp, index+i, str, usedSet)
			tmp = tmp[:len(tmp)-1]
		}
	}
}

func restoreIpAddresses(s string) []string {
	result := []string{}
	usedSet := make(map[string]bool)
	dfs(&result, []int{}, 0, s, &usedSet)
	return result
}
