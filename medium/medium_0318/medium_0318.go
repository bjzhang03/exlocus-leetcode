package medium_0318

func strMul(a string, b string) int {
	result := 0
	if len(a) > 0 && len(b) > 0 {
		// fmt.Println(a, " - ", b)
		countA := map[uint8]int{'a': 0, 'b': 0, 'c': 0, 'd': 0, 'e': 0, 'f': 0, 'g': 0, 'h': 0, 'i': 0, 'j': 0, 'k': 0, 'l': 0, 'm': 0, 'n': 0, 'o': 0, 'p': 0, 'q': 0, 'r': 0, 's': 0, 't': 0, 'u': 0, 'v': 0, 'w': 0, 'x': 0, 'y': 0, 'z': 0}
		for i := 0; i < len(a); i++ {
			countA[a[i]] = 1
		}
		for i := 0; i < len(b); i++ {
			if countA[b[i]] > 0 {
				return result
			}
		}
		result = len(a) * len(b)
	}
	return result
}

func maxProduct(words []string) int {
	result := 0
	if len(words) > 0 {
		for i := 0; i < len(words); i++ {
			for j := i + 1; j < len(words); j++ {
				tmp := 0
				if len(words[i]) > len(words[j]) {
					tmp = strMul(words[j], words[i])
				} else {
					tmp = strMul(words[i], words[j])
				}
				if result < tmp {
					result = tmp
				}
			}
		}

	}
	return result
}
