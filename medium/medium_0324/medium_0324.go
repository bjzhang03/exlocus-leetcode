package medium_0324

import (
	"sort"
)

//https://www.cnblogs.com/grandyang/p/5139057.html
func wiggleSort(nums []int) {
	if len(nums) > 0 {
		// 先对数据进行排序操作
		count := len(nums) / 2
		if len(nums)%2 != 0 {
			count = len(nums)/2 + 1
		}
		sort.Ints(nums)

		left := append([]int{}, nums[:count]...)
		right := append([]int{}, nums[count:]...)

		nums = nums[0:0]
		// 从后向前加数据,这个设计很巧妙,我没想出来，参照别人的思路
		for len(left) > 0 && len(right) > 0 {
			nums = append(nums, left[len(left)-1])
			nums = append(nums, right[len(right)-1])
			left = left[:len(left)-1]
			right = right[:len(right)-1]
		}
		if len(left) > 0 {
			nums = append(nums, left[0])
		}
	}
}
