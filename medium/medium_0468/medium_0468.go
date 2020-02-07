package medium_0468

import (
	"strconv"
	"strings"
)

func validIPAddress(IP string) string {
	if checkIpv4(IP) {
		return "IPv4"
	} else if checkIpv6(strings.ToLower(IP)) {
		return "IPv6"
	}
	return "Neither"
}

func checkIpv4(ip string) bool {
	if len(ip) > 0 {
		// 先对ip地址进行切分
		strs := strings.Split(ip, ".")
		// 对长度进行判断
		if len(strs) != 4 {
			return false
		}
		// 对每一个切分后的字符串进行判断
		for i := 0; i < len(strs); i++ {
			// 判断长度
			if !(len(strs[i]) <= 3 && len(strs[i]) > 0) {
				return false
			}
			// 判断范围
			if !func() bool {
				for j := 0; j < len(strs[i]); j++ {
					if !(strs[i][j] >= '0' && strs[i][j] <= '9') {
						return false
					}
				}
				if strnum, _ := strconv.Atoi(strs[i]); !(strnum >= 0 && strnum <= 255) {
					return false
				}
				return true
			}() {
				return false
			}
			//判断是否有前导0
			if len(strs[i]) > 1 && strs[i][0] == '0' {
				return false
			}
		}
		return true
	}
	return false
}

func checkIpv6(ip string) bool {
	if len(ip) > 0 {
		strs := strings.Split(ip, ":")
		// 检查长度
		if len(strs) != 8 {
			return false
		}
		for i := 0; i < len(strs); i++ {
			// 判断长度
			if !(len(strs[i]) <= 4 && len(strs[i]) > 0) {
				return false
			}
			// 判断数字范围
			if !func(str string) bool {
				for i := 0; i < len(str); i++ {
					flag := false
					if (str[i] >= '0' && str[i] <= '9') {
						flag = true
					}
					if (str[i] >= 'a' && str[i] <= 'f') {
						flag = true
					}
					if !flag {
						return false
					}
				}
				return true
			}(strs[i]) {
				return false
			}
		}
		return true
	}
	return false
}
