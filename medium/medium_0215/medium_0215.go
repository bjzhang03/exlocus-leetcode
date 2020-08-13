package medium_0215

func headAdjust(nums *[]int, position int) {
	if position < len(*nums) {
		//fmt.Println("position = ", position, "nums = ", *nums)
		left := position * 2
		right := position*2 + 1

		if right < len(*nums) && (*nums)[position] > (*nums)[right] {
			tmp := (*nums)[position]
			(*nums)[position] = (*nums)[right]
			(*nums)[right] = tmp
		}

		if left < len(*nums) && (*nums)[position] > (*nums)[left] {
			tmp := (*nums)[position]
			(*nums)[position] = (*nums)[left]
			(*nums)[left] = tmp
		}
		// 对堆数据进行递归调整
		headAdjust(nums, left)
		headAdjust(nums, right)
	}
}

func findKthLargest(nums []int, k int) int {
	result := 0
	if len(nums) > 0 && len(nums) >= k {
		// 新建数组
		savek := []int{0}
		for i := 0; i < len(nums); i++ {
			if len(savek) < k {
				savek = append(savek, nums[i])
			} else if len(savek) == k {
				savek = append(savek, nums[i])
				for i := k; i >= 1; i-- {
					headAdjust(&savek, i)
				}
			} else if savek[1] < nums[i] {
				savek[1] = nums[i]
			}
			headAdjust(&savek, 1)
			//fmt.Println("save = ", savek)
		}
		result = savek[1]
		//fmt.Println("save = ", savek)
	}
	return result
}
