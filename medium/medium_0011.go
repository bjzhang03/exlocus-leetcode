package medium

func maxArea(height []int) int {
	result := 0
	if len(height) > 1 {
		for i := 0; i < len(height); i++ {
			for j := i + 1; j < len(height); j++ {
				tmp := 0
				if height[i] > height[j] {
					tmp = height[j] * (j - i)
				} else {
					tmp = height[i] * (j - i)
				}
				if result < tmp {
					result = tmp
				}
			}
		}
	}
	return result

}
