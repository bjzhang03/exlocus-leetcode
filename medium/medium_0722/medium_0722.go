package medium_0722

import (
	"strings"
)

// https://www.cnblogs.com/jasminemzy/p/9739410.html
func removeComments(source []string) []string {
	if len(source) > 0 {
		result := []string{}
		stack := []string{}
		isincomment := false
		for i := 0; i < len(source); i++ {
			for j := 0; j < len(source[i]); j++ {
				if j+1 < len(source[i]) {
					ctmp := string(source[i][j]) + string(source[i][j+1])
					if !isincomment && ctmp == "//" {
						result = append(result, source[i][:j])
						break
					} else if !isincomment && ctmp == "/*" {
						isincomment = true
						stack = append(stack, source[i])
					} else if isincomment && ctmp == "*/" {
						isincomment = false
						result = append(result, stack[0][:strings.Index(stack[0], "/*")]+source[i][j+2:])
						stack = stack[0:0]
					}
				}
			}
			if !isincomment {
				result = append(result, source[i])
			}
		}
		return result
	}
	return nil
}
