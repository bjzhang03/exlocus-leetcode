package medium

import (
	"fmt"
	"strings"
)

func isValidSerialization(preorder string) bool {
	result := false
	if len(preorder) > 0 {
		strs := strings.Split(preorder, ",")
		fmt.Println(strs)

	}
	return result
}
