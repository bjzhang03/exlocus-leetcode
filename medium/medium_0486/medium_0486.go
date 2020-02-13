package medium_0486

import "fmt"

func PredictTheWinner(nums []int) bool {
	pone, ptwo := solve(nums)
	fmt.Println(pone, ptwo)
	return pone >= ptwo
}

func solve(nums []int) (int, int) {
	pone, ptwo := 0, 0
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			pone, ptwo = nums[0], nums[1]
		} else {
			pone, ptwo = nums[1], nums[0]
		}
	} else if len(nums) > 2 {
		// 当one选择第一个
		if nums[0] > nums[len(nums)-1] {
			if nums[1] > nums[len(nums)-1] {
				nnums := nums[2:]
				npone, nptwo := solve(nnums)
				pone, ptwo = npone+nums[0], nptwo+nums[1]
			} else {
				nnums := nums[1 : len(nums)-1]
				npone, nptwo := solve(nnums)
				pone, ptwo = npone+nums[0], nptwo+nums[len(nums)-1]
			}
		} else {
			// 当one选择最后一个
			if nums[0] > nums[len(nums)-2] {
				nnums := nums[1 : len(nums)-1]
				npone, nptwo := solve(nnums)
				pone, ptwo = npone+nums[len(nums)-1], nptwo+nums[0]
			} else {
				nnums := nums[:len(nums)-2]
				npone, nptwo := solve(nnums)
				pone, ptwo = npone+nums[len(nums)-1], nptwo+nums[len(nums)-2]
			}
		}
	}
	return pone, ptwo
}
