package medium_0658

import "math"

func findClosestElements(arr []int, k int, x int) []int {
	// 留下最近的,去掉最远的,直到数组大大小为k为止
	if len(arr) > 0 && k > 0 {
		for len(arr) > k {
			// 查看第一个和最后一个哪一个离当前的数字远
			if math.Abs(float64(arr[0]-x)) > math.Abs(float64(arr[len(arr)-1]-x)) {
				arr = arr[1:]
			} else {
				arr = arr[:len(arr)-1]
			}
		}
		return arr
	}
	return nil
}
