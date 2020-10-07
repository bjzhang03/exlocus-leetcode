package medium_0740

import (
	"fmt"
	"sort"
)

func deleteAndEarn(nums []int) int {
	if len(nums) > 0 {
		// 先对数据进行排序操作
		sort.Ints(nums)
		// 先对数据进行存储操作,计算出每一个数据有多少个
		save := make(map[int]int, 10000)
		// 存储去重之后的数据
		uniqnums := []int{}
		for i := 0; i < len(nums); i++ {
			uniqnums = append(uniqnums, nums[i])
			save[nums[i]] = 1
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
				save[nums[i]]++
			}
		}
		return solve(uniqnums, &save, &map[string]int{})
	}
	return 0
}

// 使用一个calculated来存储计算过的数据,以便减少计算量
func solve(uniqnums []int, save *map[int]int, calculated *map[string]int) int {
	// 如果是已经计算过的结果则直接返回即可
	if _, ok := (*calculated)[fmt.Sprint(uniqnums)]; ok {
		return (*calculated)[fmt.Sprint(uniqnums)]
	} else if !ok {
		result := 0
		if len(uniqnums) == 1 {
			// 只有一个元素
			result = (*save)[uniqnums[0]] * uniqnums[0]
		} else if len(uniqnums) > 1 && uniqnums[0]+1 != uniqnums[1] {
			// 第一个一定可以取到
			result = (*save)[uniqnums[0]]*uniqnums[0] + solve(append([]int{}, uniqnums[1:]...), save, calculated)
		} else if len(uniqnums) > 1 && uniqnums[0]+1 == uniqnums[1] {
			// 第一个元素可能没有办法取到
			if len(uniqnums) == 2 {
				// 只有两个元素的情况
				if (*save)[uniqnums[0]]*uniqnums[0] > (*save)[uniqnums[1]]*uniqnums[1] {
					result = (*save)[uniqnums[0]] * uniqnums[0]
				} else {
					result = (*save)[uniqnums[1]] * uniqnums[1]
				}
			} else if len(uniqnums) > 2 {
				// 元素多余两个的情况
				if uniqnums[1]+1 != uniqnums[2] {
					if (*save)[uniqnums[0]]*uniqnums[0] > (*save)[uniqnums[1]]*uniqnums[1] {
						result = (*save)[uniqnums[0]]*uniqnums[0] + solve(append([]int{}, uniqnums[2:]...), save, calculated)
					} else {
						result = (*save)[uniqnums[1]]*uniqnums[1] + solve(append([]int{}, uniqnums[2:]...), save, calculated)
					}
				} else if uniqnums[1]+1 == uniqnums[2] {
					selectone := (*save)[uniqnums[0]]*uniqnums[0] + solve(append([]int{}, uniqnums[2:]...), save, calculated)
					selecttwo := (*save)[uniqnums[1]]*uniqnums[1] + solve(append([]int{}, uniqnums[3:]...), save, calculated)
					if selectone > selecttwo {
						result = selectone
					} else {
						result = selecttwo
					}
				}
			}
		}
		(*calculated)[fmt.Sprint(uniqnums)] = result
		return result
	}
	return 0
}
