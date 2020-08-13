package medium_0673

func findNumberOfLIS(nums []int) int {
	if len(nums) > 0 {
		save := make([]sinfo, len(nums))
		// 首先默认长度是1
		maxlen := 1
		for i := 0; i < len(nums); i++ {
			tmp := *&sinfo{1, 1}
			for j := i - 1; j >= 0; j-- {
				// 找到前面出现的比它小的数据
				if nums[j] < nums[i] {
					// 如果长度可以变得更长,则更新数据
					if save[j].length+1 > tmp.length {
						tmp = *&sinfo{save[j].length + 1, save[j].count}
					} else if save[j].length+1 == tmp.length {
						// 如果可以变得更多,则进行统计
						tmp.count = tmp.count + save[j].count
					}
				}
			}
			// 寻找到长度最长的数据
			if maxlen < tmp.length {
				maxlen = tmp.length
			}
			// 以当前的数据为结尾的数据的最长increase的统计信息
			save[i] = tmp
		}
		// 找到所有的最大长度的数据
		result := 0
		for i := 0; i < len(save); i++ {
			if save[i].length == maxlen {
				result = result + save[i].count
			}
		}
		return result
	}
	return 0
}

type sinfo struct {
	length int
	count  int
}
