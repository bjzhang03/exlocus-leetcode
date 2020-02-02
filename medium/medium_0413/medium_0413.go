package medium_0413

func numberOfArithmeticSlices(A []int) int {
	result := 0
	if len(A) > 0 {
		for i := 0; i < len(A); i++ {
			tmps := []int{}
			for j := i; j >= 0; j-- {
				tmps = append(tmps, A[j])
				if len(tmps) >= 3 && isArithmetic(tmps) {
					result++
				} else if len(tmps) >= 3 && !isArithmetic(tmps) {
					break
				}
			}
		}
	}
	return result
}

func isArithmetic(nums []int) bool {
	result := true
	if len(nums) < 3 {
		result = false
	} else {
		curs := nums[1] - nums[0]
		for i := 2; i < len(nums); i++ {
			if nums[i]-nums[i-1] != curs {
				result = false
			}
		}
	}
	return result
}
