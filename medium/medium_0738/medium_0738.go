package medium_0738

func monotoneIncreasingDigits(N int) int {
	if N > 0 {
		// 先将大的数字存储成想要的数据
		save := []int{}
		for N > 0 {
			save = append([]int{N % 10}, save...)
			N = N / 10
		}
		// 构造数据
		for !checkMonotoneIncreasing(save) {
			// 构造需要用到的数据
			ress := []int{save[0]}
			for i := 1; i < len(save); i++ {
				if save[i] < save[i-1] {
					//根据ress构造结果数据
					tmpsum := 0
					for i := 0; i < len(ress); i++ {
						tmpsum = tmpsum*10 + ress[i]
					}
					//数字减去1,然后再来一次
					tmpsum = tmpsum - 1
					ress = ress[:0]
					for tmpsum > 0 {
						ress = append([]int{tmpsum % 10}, ress...)
						tmpsum = tmpsum / 10
					}
					// 添加一个最大的元素进来
					for j := i; j < len(save); j++ {
						ress = append(ress, 9)
					}
					// 构建完数据,并结束这个for循环,重新开始下一层的循环
					break
				} else {
					ress = append(ress, save[i])
				}
			}
			save = save[0:0]
			save = append(save, ress...)
		}
		// 构造返回的数字
		result := 0
		for i := 0; i < len(save); i++ {
			result = result*10 + save[i]
		}
		return result
	}
	return 0
}

func checkMonotoneIncreasing(nums []int) bool {
	if len(nums) > 0 {
		flag := nums[0]
		for i := 1; i < len(nums); i++ {
			if nums[i] < flag {
				return false
			}
			flag = nums[i]
		}
	}
	return true
}
