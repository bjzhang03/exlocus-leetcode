package medium_0718

// 这个题目时lcs的一个变种
func findLength(A []int, B []int) int {
	if len(A) > 0 && len(B) > 0 {
		save := make([][]int, len(A)+1)
		for i := 0; i < len(save); i++ {
			save[i] = make([]int, len(B)+1)
		}
		result := 0
		for i := 0; i < len(A); i++ {
			for j := 0; j < len(B); j++ {
				if A[i] == B[j] {
					save[i+1][j+1] = save[i][j] + 1
				}
				if result < save[i+1][j+1] {
					result = save[i+1][j+1]
				}
			}
		}
		return result
	}
	return 0
}
