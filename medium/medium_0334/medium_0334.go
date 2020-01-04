package medium_0334

func increasingTriplet(nums []int) bool {
	if len(nums) > 2 {

		for i := 0; i < len(nums); i++ {
			lagercount := 0
			currentnu := nums[i]

			for j := i + 1; j < len(nums); j++ {
				// 防止出现1,5,2,4这种情况的数据
				if nums[j] > currentnu {
					currentnu = nums[j]
					lagercount++
				} else if nums[j] > nums[i] && currentnu > nums[j] {
					currentnu = nums[j]
				}
				if lagercount >= 2 {
					return true
				}
			}
		}
	}
	return false
}
