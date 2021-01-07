package medium_0756

func pyramidTransition(bottom string, allowed []string) bool {
	return solve(bottom, allowed)

}

func solve(bottom string, allowed []string) bool {
	//fmt.Printf("solve bottom := %s, allowed := %s\n", fmt.Sprint(bottom), fmt.Sprint(allowed))
	if len(bottom) == 2 {
		for i := 0; i < len(allowed); i++ {
			if bottom == allowed[i][:2] {
				return true
			}
		}
		return false

	} else if len(bottom) > 2 {
		result := map[string][][]string{}
		deepFirstSearch(bottom, allowed, "", &result)

		for key, val := range result {
			for _, sval := range val {
				if solve(key, sval) {
					return true
				}
			}
		}
	}
	return false
}

func deepFirstSearch(bottom string, allowed []string, save string, result *map[string][][]string) {
	//fmt.Printf("bottom := %s, allowed := %s\n", bottom, fmt.Sprint(allowed))
	if len(bottom) == 2 {
		for i := 0; i < len(allowed); i++ {
			if bottom == allowed[i][:2] {
				save = save + allowed[i][2:]
				if _, ok := (*result)[save]; !ok {
					(*result)[save] = [][]string{append(append([]string{}, allowed[:i]...), allowed[i+1:]...)}
				} else if ok {
					(*result)[save] = append((*result)[save], append(append([]string{}, allowed[:i]...), allowed[i+1:]...))
				}

				save = save[:len(save)-1]
			}
		}
	} else if len(bottom) > 2 {
		current := bottom[:2]
		for i := 0; i < len(allowed); i++ {
			if current == allowed[i][:2] {
				save = save + allowed[i][2:]
				deepFirstSearch(bottom[1:], append(append([]string{}, allowed[:i]...), allowed[i+1:]...), save, result)
				save = save[:len(save)-1]
			}
		}
	}

}
