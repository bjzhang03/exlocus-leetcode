package medium_0781

func numRabbits(answers []int) int {
	result := 0
	if len(answers) > 0 {
		save := map[int]int{}
		for i := 0; i < len(answers); i++ {
			if answers[i] == 0 {
				result++
			} else if _, ok := save[answers[i]]; ok {
				save[answers[i]]++
			} else if !ok {
				save[answers[i]] = 1
			}
		}

		for key, val := range save {
			if key+1 >= val {
				result = result + key + 1
			} else if key+1 < val {
				result = result + func(a, count int) int {
					fres := count / (a + 1)
					if count%(a+1) != 0 {
						fres++
					}
					return fres
				}(key, val)*(key+1)
			}
		}
	}
	return result
}
