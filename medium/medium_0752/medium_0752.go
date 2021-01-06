package medium_0752

import "fmt"

func openLock(deadends []string, target string) int {
	return solve(func(data []string) map[string]bool {
		result := map[string]bool{}
		for i := 0; i < len(data); i++ {
			result[data[i]] = true
		}
		return result
	}(deadends), target)
}

// 这个题目,只能采用自底向上的方案,不能采用自顶向下的方案,向下的方案不好判断已经使用过的数据,容易陷入死循环
func solve(deadends map[string]bool, target string) int {
	save := []string{"0000"}
	cres := 0
	sld := map[string]bool{"0000": true}
	for len(save) > 0 {
		nsave := []string{}
		for i := 0; i < len(save); i++ {
			if save[i] == target {
				return cres
			}
			// 当前的这个数据不是 死掉的数据
			if _, okd := deadends[save[i]]; !okd {
				for j := 0; j < len(save[i]); j++ {
					// 生成新的数据
					current := save[i][j] - '0'
					fnext := save[i][:j] + fmt.Sprint((current+1)%10) + save[i][j+1:]
					bnext := save[i][:j] + fmt.Sprint((current+9)%10) + save[i][j+1:]
					// 检查数据是否已经被使用过了
					if _, ok := sld[fnext]; !ok {
						nsave = append(nsave, fnext)
						sld[fnext] = true
					}
					if _, ok := sld[bnext]; !ok {
						nsave = append(nsave, bnext)
						sld[bnext] = true
					}
				}
			}
		}
		// 到这里已经处理完一层的数据了
		cres = cres + 1
		save = append(save[:0], nsave...)
	}
	return -1
}
