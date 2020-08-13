package medium_0670

func maximumSwap(num int) int {
	if num > 0 {
		save := []int{}
		for num > 0 {
			save = append([]int{num % 10}, save...)
			num = num / 10
		}
		// 从前向后按个寻找数据
		for i := 0; i < len(save); i++ {
			current := save[i]
			flag := false
			target := current
			index := i
			// 从后面找到一个最大的数据出来
			for j := i + 1; j < len(save); j++ {
				if save[j] > target {
					target = save[j]
					index = j
					flag = true
				} else if save[j] == target && index < j {
					// 数字相等的时候则选择后面的那个
					index = j
				}
			}
			// 找到结果,则进行数据位置的交换
			if flag {
				tmp := save[i]
				save[i] = save[index]
				save[index] = tmp
				break
			}
		}
		// 构造结果
		result := 0
		for k := 0; k < len(save); k++ {
			result = 10*result + save[k]
		}
		return result
	}
	return 0
}
