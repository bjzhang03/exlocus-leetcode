package medium_0777

import "fmt"

// 这个题目我没有想出来解法
func canTransform(start string, end string) bool {
	return solve(start, end)
}

func solve(start string, end string) bool {
	fmt.Printf("solve start:= %s, end:= %s\n", start, end)
	if len(start) == len(end) {
		// 长度为0,则返回true
		if len(start) == 0 {
			return true
		} else {
			// start以R为开始,end以L为标记开始处理
			for i := 0; i < len(start); i++ {
				if start[i] == 'R' && end[i] == 'X' {
					ok, nstart, nend := handler(start[i:], end[i:])
					if ok {
						return solve(nstart, nend)
					} else if !ok {
						return false
					}
				} else if end[i] == 'L' && start[i] == 'X' {
					ok, nstart, nend := handlel(start[i:], end[i:])
					fmt.Printf("after handlel ok:= %s, nstart := %s, nend := %s\n", fmt.Sprint(ok), fmt.Sprint(nstart), fmt.Sprint(nend))
					if ok {
						return solve(nstart, nend)
					} else if !ok {
						return false
					}
				} else if !(start[i] == end[i]) {
					return false
				}
			}
			// 全部处理完了,都是一样的符号,则直接返回true
			return true
		}
	}
	return false
}

func handler(start, end string) (bool, string, string) {
	fmt.Printf("handler start:= %s, end := %s\n", start, end)
	if start[0] == 'R' && end[0] == 'X' {
		for i := 1; i < len(start); i++ {
			if start[i] == 'X' && end[i] == 'X' {
				// 啥事情都不做
			} else if start[i] == 'X' && end[i] == 'R' {
				return true, start[i+1:], end[i+1:]
			} else if !(start[i] == 'X' && end[i] == 'X') {
				return false, "", ""
			}
		}
	}
	return false, "", ""
}

func handlel(start, end string) (bool, string, string) {
	fmt.Printf("handlel start:= %s, len(start):= %s, end := %s, len(end):= %s\n", start, fmt.Sprint(len(start)), end, fmt.Sprint(len(end)))
	if end[0] == 'L' && start[0] == 'X' {
		for i := 1; i < len(end); i++ {
			if start[i] == 'X' && end[i] == 'X' {
				// 啥事情都不做
			} else if start[i] == 'L' && end[i] == 'X' {
				return true, start[i+1:], end[i+1:]
			} else if !(start[i] == 'X' && end[i] == 'X') {
				return false, "", ""
			}
		}
	}
	return false, "", ""
}
