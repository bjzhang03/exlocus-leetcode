package medium_0473

import "sort"

// 计算sum的函数变量
var calculatesum = func(numbers []int) int {
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	return sum
}

// 思路是正确的,但是一直超时,我也不知道是啥原因
// 这道题目还是没有做出来
// https://blog.csdn.net/qq_40803710/article/details/80274628
func makesquare(nums []int) bool {
	if checkPass(nums) {
		// 先排序,将大的数据放在前面,可以很大程度上改善算法
		sort.Ints(nums)
		reversenums := []int{}
		for i := len(nums) - 1; i >= 0; i-- {
			reversenums = append(reversenums, nums[i])
		}
		return deepFirstSearch(&reversenums, []int{}, []int{}, []int{}, []int{}, calculatesum(nums)/4, 0)
	}
	return false
}

func deepFirstSearch(nums *[]int, one []int, two []int, three []int, four []int, target int, index int) bool {
	//fmt.Println(len(nums), nums, one, two, three, four, target)
	// 判断当前的状态是否可以checkpass
	result := false
	if !func(on []int, tw []int, th []int, fo []int, ta int) bool {
		//fmt.Println("overflow")
		// 对每一个数组的sum进行判断
		if calculatesum(on) > ta || calculatesum(tw) > ta || calculatesum(th) > ta || calculatesum(fo) > ta {
			return false
		}
		return true
	}(one, two, three, four, target) {
		result = false
	} else if func(on []int, tw []int, th []int, fo []int, ta int) bool {
		//fmt.Println("success")
		// 判断现在的数据是否已经是满足需求了
		result := true
		if !(calculatesum(on) == ta && calculatesum(tw) == ta && calculatesum(th) == ta && calculatesum(fo) == ta) {
			result = false
		}
		return result
	}(one, two, three, four, target) {
		result = true
	} else if len(*nums) > index {
		// 取出第一个数字即可
		current := (*nums)[index]
		// 下一次可选的数字
		// 对每一个数字进行添加测试
		// 对每一条边进行测试
		tmpone := append(one, current)
		if calculatesum(tmpone) <= target && deepFirstSearch(nums, tmpone, append([]int{}, two...), append([]int{}, three...), append([]int{}, four...), target, index+1) {
			return true
		}
		tmptwo := append(two, current)
		if calculatesum(tmptwo) <= target && deepFirstSearch(nums, append([]int{}, one...), tmptwo, append([]int{}, three...), append([]int{}, four...), target, index+1) {
			return true
		}
		tmpthree := append(three, current)
		if calculatesum(tmpthree) <= target && deepFirstSearch(nums, append([]int{}, one...), append([]int{}, two...), tmpthree, append([]int{}, four...), target, index+1) {
			return true
		}
		tmpfour := append(four, current)
		if calculatesum(tmpfour) <= target && deepFirstSearch(nums, append([]int{}, one...), append([]int{}, two...), append([]int{}, three...), tmpfour, target, index+1) {
			return true
		}
	}
	return result
}

// 先做一些基础的检查
func checkPass(nums []int) bool {
	if !func(n []int) bool {
		// 火柴的根数不能小于4
		if len(n) < 4 {
			return false
		}
		sum := 0
		// 判断是不是每一个数字都是非负数字
		for i := 0; i < len(n); i++ {
			if n[i] < 0 {
				return false
			}
			sum = sum + n[i]
		}
		// 判断总和是不是4的倍数
		if sum%4 != 0 {
			return false
		}
		// 判断是否出现有一条火柴的长度大于正方形的边长这种情况
		edge := sum / 4
		for i := 0; i < len(n); i++ {
			if n[i] > edge {
				return false
			}
		}
		return true
	}(nums) {
		return false
	}
	return true
}
