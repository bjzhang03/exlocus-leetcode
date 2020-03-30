package medium_0713

// https: //blog.csdn.net/weixin_37373020/article/details/80959169
// 这个题目我没想明白为啥是这么做的
// 我对这个题目的思考力度不够啊
func numSubarrayProductLessThanK(nums []int, k int) int {
	if len(nums) > 0 {
		count, left, one := 0, 0, 1
		for right := 0; right < len(nums); right++ {
			one *= nums[right];
			for one >= k && left <= right {
				// 这里为啥就没有出现nums[left]为0的问题?
				one /= nums[left];
				left++;
			}
			count += right - left + 1;
		}
		return count
	}
	return 0
}
