package medium_0454

func fourSumCount(A []int, B []int, C []int, D []int) int {
	result := 0
	if len(A) > 0 && len(B) > 0 && len(C) > 0 && len(D) > 0 {
		// 先构造map
		cdsumMap := func(C []int, D []int) map[int]int {
			result := map[int]int{}
			for i := 0; i < len(C); i++ {
				for j := 0; j < len(D); j++ {
					if _, ok := result[C[i]+D[j]]; ok {
						result[C[i]+D[j]]++
					} else if !ok {
						result[C[i]+D[j]] = 1
					}
				}
			}
			return result
		}(C, D)
		// 计算结果数据

		for i := 0; i < len(A); i++ {
			for j := 0; j < len(B); j++ {
				if _, ok := cdsumMap[0-A[i]-B[j]]; ok {
					result = result + cdsumMap[0-A[i]-B[j]]
				}
			}
		}
	}

	return result
}
