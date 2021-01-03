package hard_0060

import (
	"fmt"
	"math"
	"strconv"
)

var permutations = []int{1, 2, 6, 24, 120, 720, 5040, 40320, 362880}

// 处理数据刚好是全变换的情况

func isAllPermutate(num []int, count int) bool {
	if permutations[len(num)-1] == count || count == 0 {
		return true
	}
	return false
}

func permutate(num []int, count int) string {

	fmt.Println(num, count)
	result := ""

	// 如果是0的话,则刚好是逆序的操作
	if isAllPermutate(num, count) {
		for j := len(num) - 1; j >= 0; j-- {
			result = result + strconv.Itoa(num[j])
		}
	} else {
		// 找到当前的需要放在第一个位置的数字
		for i := len(num) - 1; i > 0; i-- {
			if permutations[i] <= count {
				// 找到下标之后,将数据添加进来
				target := (count / permutations[i])
				// 判断一下是否是全变换,这里的数据已经用完的全部变换
				if int(math.Mod(float64(count), float64(permutations[i]))) == 0 {
					target = target - 1
				}
				result = result + strconv.Itoa(num[target])
				// 构建新的候选数据集
				nextNum := num[:target]
				nextNum = append(nextNum, num[target+1:]...)
				// 循环添加进来数据
				result = result + permutate(nextNum, int(math.Mod(float64(count), float64(permutations[i]))))
				return result
			}
		}
		// 如果找到的是第一个排列,则直接顺序列出即可
		for j := 0; j < len(num); j++ {
			result = result + strconv.Itoa(num[j])
		}
	}
	return result
}

func getPermutation(n int, k int) string {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	return permutate(nums[:n], k)
}
