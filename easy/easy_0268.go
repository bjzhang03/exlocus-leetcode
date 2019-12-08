package easy

func missingNumber(nums []int) int {
	result := -1
	if len(nums) > 0 {
		nums = append(nums, -1)
		for i := 0; i < len(nums); i++ {
			for nums[i] != i {
				if nums[i] == -1 {
					flag := true
					for j := i; j < len(nums); j++ {
						if nums[j] == i {
							nums[i] = i
							nums[j] = -1
							flag = false
							break
						}
					}
					if flag {
						result = i
						break
					}
				} else {
					tmp := nums[i]
					nums[i] = nums[nums[i]]
					nums[tmp] = tmp
				}
			}
		}
	}
	return result
}
