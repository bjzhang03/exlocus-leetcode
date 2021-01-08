package medium_0769

import "math"

func maxChunksToSorted(arr []int) int {
	return solve(arr)
}

func solve(arr []int) int {
	result := 1
	for i := 0; i < len(arr); i++ {
		if sliceCompare(arr[:i+1], arr[i+1:]) {
			return result + solve(arr[i+1:])
		}
	}
	return result
}

func sliceCompare(left, right []int) bool {
	if len(left) > 0 && len(right) > 0 {
		//left找到最大的
		lmax := math.MinInt32
		for i := 0; i < len(left); i++ {
			if lmax < left[i] {
				lmax = left[i]
			}
		}
		rmin := math.MaxInt32
		for i := 0; i < len(right); i++ {
			if rmin > right[i] {
				rmin = right[i]
			}
		}
		return lmax < rmin
	}
	return false
}
