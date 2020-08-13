package medium_0331

import (
	"strings"
)

func isValidSerialization(preorder string) bool {
	if len(preorder) > 0 {
		orderstr := strings.Split(preorder, ",")
		for len(orderstr) > 0 {
			if orderstr[0] == "#" {
				break
			}
			next := []string{}
			for i := 0; i < len(orderstr); i++ {
				if orderstr[i] != "#" && i+2 < len(orderstr) && orderstr[i+1] == "#" && orderstr[i+2] == "#" {
					next = append(next, orderstr[:i]...)
					next = append(next, "#")
					if i+3 < len(orderstr) {
						next = append(next, orderstr[i+3:]...)
					}
					break
				}
			}
			orderstr = next
		}
		return len(orderstr) == 1
	}
	return true
}
