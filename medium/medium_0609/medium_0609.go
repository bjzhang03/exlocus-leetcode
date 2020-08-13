package medium_0609

import "strings"

func findDuplicate(paths []string) [][]string {
	if len(paths) > 0 {
		save := map[string][]string{}
		for i := 0; i < len(paths); i++ {
			tmp := strToMap(paths[i])
			for key, val := range tmp {
				if _, ok := save[key]; ok {
					save[key] = append(save[key], val...)
				} else if !ok {
					save[key] = append([]string{}, val...)
				}
			}
		}

		result := [][]string{}
		for _, val := range save {
			if len(val) > 1 {
				result = append(result, val)
			}
		}
		return result
	}
	return nil
}

func strToMap(path string) map[string][]string {
	pathSlice := strings.Split(path, " ")
	dir := pathSlice[0]
	save := map[string][]string{}
	for i := 1; i < len(pathSlice); i++ {
		left := func(s string) int {
			for j := len(s) - 2; j >= 0; j-- {
				if s[j] == '(' {
					return j
				}
			}
			return 0
		}(pathSlice[i])
		content := pathSlice[i][left+1 : len(pathSlice[i])-1]
		name := pathSlice[i][:left]
		if _, ok := save[content]; ok {
			save[content] = append(save[content], dir+"/"+name)
		} else if !ok {
			save[content] = append([]string{}, dir+"/"+name)
		}
	}
	return save
}
