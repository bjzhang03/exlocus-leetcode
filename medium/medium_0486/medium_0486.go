package medium_0486

import "fmt"

// 这个题目是典型的minmax算法,不过我之前没有处理过minmax算法,所以我不会做,难过ing
// 这里除了进行minmax算法之后,还进行了alpha-beta剪枝,aplha-beta剪枝 也是一个很好的思想
// https://blog.csdn.net/zkybeck_ck/article/details/45644471
// https://blog.csdn.net/pdfcxc/article/details/89281249
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
