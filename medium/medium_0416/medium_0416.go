package medium_0416

func canPartition(nums []int) bool {
	if len(nums) > 0 {
		csum := 0
		for i := 0; i < len(nums); i++ {
			csum = csum + nums[i]
		}
		if csum%2 != 0 {
			return false
		} else {
			// 这里使用的是子集和问题的解法,子集和问题是np问题
			numsmap := map[int]int{}
			for i := 0; i < len(nums); i++ {
				if _, ok := numsmap[nums[i]]; !ok {
					numsmap[nums[i]] = 1
				} else {
					numsmap[nums[i]]++
				}
			}
			save := make([]map[int]int, csum/2+1)
			// 只需要计算到一半即可,不需要全部计算完
			for i := 1; i <= csum/2; i++ {
				if _, ok := numsmap[i]; ok {
					// 数据直接存在的话,则直接添加进来即可
					save[i] = map[int]int{i: 1}
				} else {
					for j := 0; j < len(nums); j++ {
						if i-nums[j] >= 0 &&
							len(save[i-nums[j]]) > 0 &&
							func(amap *map[int]int, bmap *map[int]int, item int) bool {
								// 使用匿名函数直接计算结果
								// map的大小不大于100
								if _, ok := (*bmap)[item]; ok && (*amap)[item] > (*bmap)[item] {
									return true

								} else if !ok {
									return true
								}
								return false
								// 函数的值使用指针,防止复制操作
							}(&numsmap, &save[i-nums[j]], nums[j]) {
							// 记录下当前使用到的数据
							// 防止动态扩容,指定map的容量
							// 这里已经明确说明数字是小于等于100的
							save[i] = make(map[int]int, 100)
							for key, val := range save[i-nums[j]] {
								save[i][key] = val
							}
							if _, ok := save[i][nums[j]]; !ok {
								save[i][nums[j]] = 1
							} else {
								save[i][nums[j]]++
							}
							// 找到一个即可终止
							// 这里是加速的核心
							break
						}
					}
				}

			}
			return len(save[csum/2]) > 0
		}
	}
	return true
}
