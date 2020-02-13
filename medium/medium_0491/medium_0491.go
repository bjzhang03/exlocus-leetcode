package medium_0491

import (
	"strconv"
)

func findSubsequences(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		deepFirstSearch(nums, &result, map[string]bool{}, []int{})
	}
	return result
}

func deepFirstSearch(nums []int, result *[][]int, used map[string]bool, tmp []int) {
	if _, ok := used[sliceToString(tmp)]; !ok && checkPass(tmp) {
		*result = append(*result, append([]int{}, tmp...))
		used[sliceToString(tmp)] = true
	}
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			// 这里还可以剪枝
			deepFirstSearch(nums[i+1:], result, used, append(tmp, nums[i]))
		}
	}
}

// 将slice转化为string
func sliceToString(sl []int) string {
	result := "("
	if len(sl) > 0 {
		for i := 0; i < len(sl)-1; i++ {
			result = result + strconv.Itoa(sl[i]) + ","
		}
		result = result + strconv.Itoa(sl[len(sl)-1])
	}
	result = result + ")"
	return result
}

func checkPass(nums []int) bool {
	result := true
	if len(nums) < 2 {
		result = false
	} else if len(nums) >= 2 {
		for i := 1; i < len(nums); i++ {
			if nums[i] < nums[i-1] {
				result = false
				break
			}
		}
	}
	return result
}
