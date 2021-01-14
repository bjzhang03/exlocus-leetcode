package medium_0784

import (
	"fmt"
	"strings"
)

func letterCasePermutation(S string) []string {
	return func(str string) []string {
		result := []string{}
		solve(str, "", &result, len(str))
		return result
	}(strings.ToLower(S))
}

func solve(str string, tmp string, result *[]string, target int) {
	if len(str) == 0 || func(data string) bool {
		for i := 0; i < len(data); i++ {
			if !(data[i] >= '0' && data[i] <= '9') {
				return false
			}
		}
		return true
	}(str) {
		if len(tmp+str) == target {
			(*result) = append(*result, tmp+str)
		}
	} else if len(str) > 0 {
		for i := 0; i < len(str); i++ {
			if str[i] >= 'a' && str[i] <= 'z' {
				tmp = tmp + fmt.Sprintf("%c", str[i])
				solve(str[i+1:], tmp, result, target)
				tmp = tmp[:len(tmp)-1]
				tmp = tmp + fmt.Sprintf("%c", str[i]-'a'+'A')
				solve(str[i+1:], tmp, result, target)
				tmp = tmp[:len(tmp)-1]
			} else {
				tmp = tmp + fmt.Sprintf("%c", str[i])
			}
		}

	}
}
