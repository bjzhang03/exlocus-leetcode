package medium_0031

import (
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	if len(nums) > 0 {
		p := 0
		for i := len(nums) - 1; i > 0; i-- {
			if nums[i] > nums[i-1] {
				p = i
				break
			}
		}
		if p > 0 {
			left := nums[:p]
			// 使用tmp来存储右边的最小的数字
			mnumber := math.MaxInt32
			mindex := -1
			right := nums[p:]
			// 在右边找到最小的数据
			for i := 0; i < len(right); i++ {
				if right[i] > left[len(left)-1] && right[i] < mnumber {
					mnumber = right[i]
					mindex = i
				}
			}

			p := right[mindex]
			right[mindex] = left[len(left)-1]
			left[len(left)-1] = p

			sort.Ints(right)
			nums = append(left, right...)
		} else {
			start := 0
			end := len(nums) - 1
			for start <= end {
				tmp := nums[end]
				nums[end] = nums[start]
				nums[start] = tmp
				end--
				start++
			}
		}
	}
}
