package medium_0015

import (
	"sort"
)

func hasslice(save [][]int, item []int) bool {
	result := false
	for i := 0; i < len(save); i++ {
		if save[i][0] == item[0] && save[i][1] == item[1] && save[i][2] == item[2] {
			result = true
			break
		}
	}
	return result
}
func validate(save map[int]int, item []int) bool {
	result := true
	tmp := make(map[int]int)
	for i := 0; i < len(item); i++ {
		if _, ok := tmp[item[i]]; ok {
			tmp[item[i]]++
		} else {
			tmp[item[i]] = 1
		}
	}

	for k, v := range tmp {
		if save[k] < v {
			result = false
			break
		}
	}
	return result
}

func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) > 0 {
		// 排序操作
		sort.Ints(nums)
		save := make(map[int]int)
		for i := 0; i < len(nums); i++ {
			if _, ok := save[nums[i]]; !ok {
				save[nums[i]] = 1
			} else {
				save[nums[i]]++
			}
		}
		for i := 0; i < len(nums); i++ {
			// 三个数都大于0的话，则不用进行查找了
			if nums[i] > 0 {
				break
			}
			// 向后面查找第二个数据
			for j := i + 1; j < len(nums); j++ {
				notsum := 0 - nums[i] - nums[j]
				if _, ok := save[notsum]; ok && notsum >= nums[j] {
					tmp := []int{nums[i], nums[j], notsum}
					if !hasslice(result, tmp) && validate(save, tmp) {
						result = append(result, tmp)
					}
				}
				for j+1 < len(nums) && nums[j] == nums[j+1] {
					j++
				}
			}
			for i+1 < len(nums) && nums[i] == nums[i+1] {
				i++
			}
		}
	}
	return result

}
