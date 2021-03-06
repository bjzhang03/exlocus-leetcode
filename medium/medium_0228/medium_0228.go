package medium_0228

import (
	"strconv"
)

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) > 0 {
		for i := 0; i < len(nums); i++ {
			tmpstr := "" + strconv.Itoa(nums[i])
			if i+1 < len(nums) && nums[i]+1 == nums[i+1] {
				tmpstr = tmpstr + "->"
				for i+1 < len(nums) && nums[i]+1 == nums[i+1] {
					i = i + 1
				}
				tmpstr = tmpstr + strconv.Itoa(nums[i])
			}
			result = append(result, tmpstr)
		}
	}
	return result
}
