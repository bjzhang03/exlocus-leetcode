package medium_0592

import (
	"strconv"
	"strings"
)

func fractionAddition(expression string) string {
	if len(expression) > 0 {
		flags := []string{}
		molecules := []int{}
		denominators := []int{}
		// 对表达式进行拆分
		for len(expression) > 0 {
			flag, molecule, denominator, nexpression := getFraction(expression)
			expression = nexpression
			flags = append(flags, flag)
			molecules = append(molecules, molecule)
			denominators = append(denominators, denominator)
		}
		// 不求最大公约数了,直接求乘积,应该也是可以的
		cDenominator := 1
		for i := 0; i < len(denominators); i++ {
			cDenominator = cDenominator * denominators[i]
		}
		// 计算分子数据
		cMolecule := 0
		for i := 0; i < len(molecules); i++ {
			if flags[i] == "+" {
				cMolecule = cMolecule + molecules[i]*(cDenominator/denominators[i])
			} else if flags[i] == "-" {
				cMolecule = cMolecule - molecules[i]*(cDenominator/denominators[i])
			}
		}
		// 拼接结果
		result := ""
		if cMolecule < 0 {
			result = "-"
			cMolecule = -cMolecule
		}
		return result + strconv.Itoa(cMolecule/getGCD(cDenominator, cMolecule)) + "/" + strconv.Itoa(cDenominator/getGCD(cDenominator, cMolecule))
	}
	return ""
}

// 获取第一个位置的分数
func getFraction(expression string) (string, int, int, string) {
	if len(expression) > 0 {
		flag := "+"
		// 判断当前的符号
		sta := 0
		if expression[sta] == '-' {
			flag = "-"
			sta++
		}
		// 找到第一个字符串出现的位置
		subp := strings.Index(expression, "/")
		molecule, _ := strconv.Atoi(expression[sta:subp])
		// 找到另一个开始的位置
		sta = subp + 1
		// 数字的长度至少为1
		end := func(st int, ex string) int {
			for i := sta; i < len(ex); i++ {
				if !(ex[i] >= '0' && ex[i] <= '9') {
					return i
				}
			}
			return len(ex)
		}(sta, expression)
		// 如果是最后一个则直接返回空白字符
		if end == len(expression) {
			denominator, _ := strconv.Atoi(expression[sta:])
			return flag, molecule, denominator, ""
		}
		// 提取分母
		denominator, _ := strconv.Atoi(expression[sta:end])
		return flag, molecule, denominator, expression[end:]
	}
	return "", 0, 0, ""
}

// 获取最大公约数的函数
func getGCD(a, b int) int {
	if a >= b {
		if b == 0 {
			return a
		} else if a%b == 0 {
			return b
		} else {
			return getGCD(b, a%b)
		}
	} else {
		return getGCD(b, a)
	}
}
