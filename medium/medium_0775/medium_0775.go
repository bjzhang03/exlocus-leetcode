package medium_0775

func isIdealPermutation(A []int) bool {
	return solveisIdealPermutation(A)
}

func solveisIdealPermutation(A []int) bool {
	if len(A) > 0 {
		for i := 0; i < len(A); i++ {
			for j := i + 1; j < len(A); j++ {
				if A[j] < A[i] && j != i+1 {
					return false
				}
			}
		}
	}
	return true
}
