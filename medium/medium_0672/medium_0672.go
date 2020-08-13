package medium_0672

import "fmt"

func flipLights(n int, m int) int {
	if n > 0 && m > 0 {
		// 构造原始的所有的都是on的数据
		lights := []bool{}
		for i := 0; i < n; i++ {
			lights = append(lights, true)
		}
		save := [][]bool{lights}

		result := solve(save, func(a int) int {
			// 6 是1,2,3的倍数,当数据远远大于6的时候，种类的数据并不会增加很多
			// 这里把数据控制在一定范围内解决这个问题
			if a > 12 {
				return a%6 + 6
			}
			return a
		}(m))
		return len(result)
	} else if n > 0 && m == 0 {
		return 1
	}
	return 0
}

func solve(save [][]bool, m int) [][]bool {
	if m > 0 {
		csave := solve(save, m-1)
		used := map[string]bool{}
		result := [][]bool{}
		for i := 0; i < len(csave); i++ {
			// 对当前的数据实行一次变换
			current := csave[i]
			// 把所有的数据变成相反的
			all := allOperation(current)
			if _, ok := used[fmt.Sprint(all)]; !ok {
				result = append(result, all)
				used[fmt.Sprint(all)] = true
			}
			// 把全部的偶数的变成相反的
			even := evenOperation(current)
			if _, ok := used[fmt.Sprint(even)]; !ok {
				result = append(result, even)
				used[fmt.Sprint(even)] = true
			}

			// 把全部的奇数
			odd := oddOperation(current)
			if _, ok := used[fmt.Sprint(odd)]; !ok {
				result = append(result, odd)
				used[fmt.Sprint(odd)] = true
			}

			// 把3k+1变成相反的
			three := threeOperation(current)
			if _, ok := used[fmt.Sprint(three)]; !ok {
				result = append(result, three)
				used[fmt.Sprint(three)] = true
			}
		}
		return result
	}
	return save
}

// 所有的数据都进行全变换
func allOperation(save []bool) []bool {
	result := append([]bool{}, save...)
	for i := 0; i < len(result); i++ {
		result[i] = !result[i]
	}
	return result
}

// 偶数的进行变换
func evenOperation(save []bool) []bool {
	result := append([]bool{}, save...)
	for i := 0; i < len(result); i = i + 2 {
		result[i] = !result[i]
	}
	return result
}

// 奇数的进行变换
func oddOperation(save []bool) []bool {
	result := append([]bool{}, save...)
	for i := 1; i < len(result); i = i + 2 {
		result[i] = !result[i]
	}
	return result
}

// 三个变换一次
func threeOperation(save []bool) []bool {
	result := append([]bool{}, save...)
	for i := 0; i < len(result); i = i + 3 {
		result[i] = !result[i]
	}
	return result
}
