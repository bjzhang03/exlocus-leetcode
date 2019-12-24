package medium_0011

// https://blog.csdn.net/m0_37402140/article/details/80370794
// 双指针的问题
// 递归解决方案
func maxArea(height []int) int {
	result := 0
	if len(height) > 1 {
		tmp := 0
		if height[0] < height[len(height)-1] {
			tmp = height[0] * (len(height) - 1)
			if result < tmp {
				result = tmp
			}
			tmpleft := maxArea(height[1:])
			if result < tmpleft {
				result = tmpleft
			}
		} else {
			tmp = height[len(height)-1] * (len(height) - 1)
			if result < tmp {
				result = tmp
			}
			tmpright := maxArea(height[:len(height)-1])
			if result < tmpright {
				result = tmpright
			}
		}
	}
	return result
}
