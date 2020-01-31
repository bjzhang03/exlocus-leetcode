package medium_0378

// 虽然ac了,但是不是最优的方案
func kthSmallest(matrix [][]int, k int) int {
	save := []int{}
	for i := 0; i < len(matrix); i++ {
		save = merge(save, matrix[i])
	}
	return save[k-1]
}

// 有序数字合并算法
func merge(save []int, nums []int) []int {
	result := []int{}

	savei := 0
	numsi := 0
	for savei < len(save) && numsi < len(nums) {

		if save[savei] < nums[numsi] {
			result = append(result, save[savei])
			savei++
		} else {
			result = append(result, nums[numsi])
			numsi++
		}
	}

	for savei < len(save) {
		result = append(result, save[savei])
		savei++
	}

	for numsi < len(nums) {
		result = append(result, nums[numsi])
		numsi++
	}
	return result

}
