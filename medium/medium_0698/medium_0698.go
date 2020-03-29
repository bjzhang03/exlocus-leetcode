package medium_0698

import (
	"sort"
)

// 求和的函数
var slicesum = func(args ... int) int {
	result := 0
	for i := 0; i < len(args); i++ {
		result = result + args[i]
	}
	return result
}

// 这里的数字不是很大,否则的话不能使用dfs进行查找,因为一定会超时
func canPartitionKSubsets(nums []int, k int) bool {
	if len(nums) > 0 && k > 0 {
		// 先计算和
		numsum := 0
		for i := 0; i < len(nums); i++ {
			numsum = numsum + nums[i]
		}

		if numsum%k == 0 {
			cmax := numsum / k
			save := make([][]int, k)
			// 排序操作
			sort.Ints(nums)
			return deepFirstSearch(nums, &save, cmax)
		} else {
			// 如果数据的和不能被k整除,则一定是false
			return false
		}
	}
	return true
}

// 深度优先遍历
func deepFirstSearch(nums []int, save *[][]int, cmax int) bool {
	if len(nums) == 0 {
		for i := 0; i < len(*save); i++ {
			if slicesum((*save)[i]...) != cmax {
				return false
			}
		}
		return true
	} else if len(nums) > 0 {
		for i := 0; i < len(*save); i++ {
			// 数据已经排序了,所以可以先从大的数据开始寻找,这样更容易找到答案
			(*save)[i] = append((*save)[i], nums[len(nums)-1])
			if slicesum((*save)[i]...) <= cmax {
				if deepFirstSearch(nums[:len(nums)-1], save, cmax) {
					return true
				}
			}
			// 还原线程
			(*save)[i] = (*save)[i][:len((*save)[i])-1]
		}
	}
	return false
}
