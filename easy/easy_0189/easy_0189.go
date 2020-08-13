package easy_0189

func rotateOneStep(nums []int) {
	tmp := nums[len(nums)-1]
	for i := len(nums) - 1; i > 0; i-- {
		nums[i] = nums[i-1]
	}
	nums[0] = tmp
}

func rotate(nums []int, k int) {
	if len(nums) > 0 && k > 0 {
		for i := 0; i < k; i++ {
			rotateOneStep(nums)
		}
	}

}
