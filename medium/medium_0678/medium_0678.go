package medium_0678

func checkValidString(s string) bool {
	if len(s) > 0 {
		return solve([]uint8{}, s)
	}
	return true
}

func solve(left []uint8, s string) bool {
	if len(left) == 0 && len(s) == 0 {
		return true
	} else if len(s) > 0 {
		for i := 0; i < len(s); i++ {
			if s[i] == '(' {
				left = append(left, '(')
			} else if s[i] == ')' {
				if len(left) > 0 {
					// 出现右括号 则直接递归操作
					left = left[:len(left)-1]
					if len(left) == 0 && len(s[i+1:]) == 0 {
						return true
					}
				} else {
					return false
				}
			} else if s[i] == '*' {
				//三种情况处理
				if len(left) > 0 {
					// 左括号,右括号,空白都可以
					return solve(append(append([]uint8{}, left...), '('), s[i+1:]) || solve(left[:len(left)-1], s[i+1:]) || solve(append([]uint8{}, left...), s[i+1:])
				} else {
					// 只能是左括号或者空白
					return solve(append(append([]uint8{}, left...), '('), s[i+1:]) || solve(append([]uint8{}, left...), s[i+1:])
				}
			}
		}
	}
	return false
}
