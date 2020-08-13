package medium_0537

import (
	"strconv"
	"strings"
)

func complexNumberMultiply(a string, b string) string {
	result := "0+0i"
	if len(a) > 0 && len(b) > 0 {
		result = strconv.Itoa(getReal(a)*getReal(b)-getComplex(b)*getComplex(a)) + "+" + strconv.Itoa(getReal(a)*getComplex(b)+getComplex(a)*getReal(b)) + "i"
	}
	return result
}

// 获取复数的实部
func getReal(s string) int {
	return func(str string) int {
		result := 0
		if i := strings.Index(str, "+"); i >= 0 {
			result, _ = strconv.Atoi(s[:i])
		}
		return result
	}(s)
}

// 获取复数的虚部
func getComplex(s string) int {
	return func(str string) int {
		result := 0
		if i := strings.Index(str, "+"); i >= 0 {
			result, _ = strconv.Atoi(s[i+1 : len(s)-1])
		}
		return result
	}(s)
}
