package easy_0167

// https://www.cnblogs.com/Harley-Quinn/p/5836149.html
// 解答链接
func twoSum(numbers []int, target int) []int {
	result := []int{0, 0}
	if len(numbers) > 0 {
		low := 0
		heigh := len(numbers) - 1
		for low < heigh {
			mid := (low + heigh) / 2
			sum := numbers[low] + numbers[heigh]
			if sum == target {
				result[0] = low + 1
				result[1] = heigh + 1
				break
			} else if sum > target {
				if numbers[low]+numbers[mid] > target {
					heigh = mid
				} else {
					heigh = heigh - 1
				}
			} else {
				if numbers[mid]+numbers[heigh] <= target {
					low = mid
				} else {
					low = low + 1
				}
			}
		}
	}

	return result
}
