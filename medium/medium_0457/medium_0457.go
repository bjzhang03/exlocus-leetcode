package medium_0457

func circularArrayLoop(nums []int) bool {
	// 只有一个数据的时候是没办法走的,因此也就谈不上循环
	if len(nums) > 1 {
		// 数据预处理,使得每个数据都不大于len(nums)
		for i := 0; i < len(nums); i++ {
			nums[i] = nums[i] % len(nums)
		}
		// 如果是全部都一样的数据,则直接返回true
		if checkEqual(nums) {
			return true
		}
		for i := 0; i < len(nums); i++ {
			// 先对数据进行预处理,保证数据不会大于nums的长度
			if nums[i] > 0 {
				if checkForward(i, nums) {
					return true
				}
			} else if nums[i] < 0 {
				if checkBack(i, nums) {
					return true
				}
			}

		}
	}
	return false
}

func checkForward(sti int, nums []int) bool {
	if len(nums) > 0 && sti >= 0 && sti < len(nums) {
		tmp := sti
		used := map[int]bool{}
		for len(used) < len(nums) {
			if nums[tmp] <= 0 {
				return false
			} else if _, ok := used[tmp]; ok {
				return true
			}
			used[tmp] = true
			// 确保数据的index正确
			tmp = (tmp + nums[tmp]) % len(nums)
		}
	}
	return false
}

func checkBack(sti int, nums []int) bool {
	if len(nums) > 0 && sti >= 0 && sti < len(nums) {
		tmp := sti
		used := map[int]bool{}
		for len(used) < len(nums) {
			if nums[tmp] >= 0 {
				return false
			} else if _, ok := used[tmp]; ok {
				return true
			}
			used[tmp] = true
			// 如果下一个数据还是自身,则应该返回false
			if tmp == (tmp+nums[tmp]+len(nums))%len(nums) {
				return false
			}
			tmp = (tmp + nums[tmp] + len(nums)) % len(nums)
		}
	}
	return false
}

// 把所有的数据全部用掉的循环,必定是所有的数据都一样的数组
func checkEqual(slice []int) bool {
	if len(slice) > 0 {
		tmp := slice[0]
		for i := 1; i < len(slice); i++ {
			if tmp != slice[i] {
				return false
			}
		}
		return true
	}
	return false
}
